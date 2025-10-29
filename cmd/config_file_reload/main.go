package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// Config 配置结构
type Config struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	LogLevel string `json:"log_level"`
	// 添加其他配置项
}

// ConfigManager 配置管理器
type ConfigManager struct {
	mu       sync.RWMutex
	config   *Config
	filePath string
}

// NewConfigManager 创建配置管理器
func NewConfigManager(filePath string) (*ConfigManager, error) {
	cm := &ConfigManager{
		filePath: filePath,
	}

	if err := cm.Reload(); err != nil {
		return nil, fmt.Errorf("初始化配置失败: %w", err)
	}

	return cm, nil
}

// Reload 重新加载配置文件
func (cm *ConfigManager) Reload() error {
	data, err := os.ReadFile(cm.filePath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	var newConfig Config
	if err := json.Unmarshal(data, &newConfig); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 可以在这里添加配置验证逻辑
	if err := cm.validateConfig(&newConfig); err != nil {
		return fmt.Errorf("配置验证失败: %w", err)
	}

	cm.mu.Lock()
	cm.config = &newConfig
	cm.mu.Unlock()

	log.Printf("配置已重新加载: %+v", newConfig)
	return nil
}

// validateConfig 验证配置有效性
func (cm *ConfigManager) validateConfig(cfg *Config) error {
	if cfg.Port <= 0 || cfg.Port > 65535 {
		return fmt.Errorf("端口号无效: %d", cfg.Port)
	}
	if cfg.Host == "" {
		return fmt.Errorf("主机地址不能为空")
	}
	return nil
}

// GetConfig 获取当前配置（返回副本）
func (cm *ConfigManager) GetConfig() Config {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return *cm.config
}

// ReloadHandler HTTP 重载接口处理函数
func (cm *ConfigManager) ReloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只支持 POST 方法", http.StatusMethodNotAllowed)
		return
	}

	log.Println("收到配置重载请求")

	if err := cm.Reload(); err != nil {
		log.Printf("重载配置失败: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "配置已重载",
		"config":  cm.GetConfig(),
		"time":    time.Now().Format(time.RFC3339),
	})
}

// GetConfigHandler 获取当前配置的接口
func (cm *ConfigManager) GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "只支持 GET 方法", http.StatusMethodNotAllowed)
		return
	}

	config := cm.GetConfig()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"config": config,
	})
}

func main() {
	// 初始化配置管理器
	configManager, err := NewConfigManager("example/need_reload.json")
	if err != nil {
		log.Fatalf("初始化配置管理器失败: %v", err)
	}

	// 注册路由
	http.HandleFunc("/api/reload", configManager.ReloadHandler)
	http.HandleFunc("/api/config", configManager.GetConfigHandler)

	// 示例业务接口，使用配置
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		config := configManager.GetConfig()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message":   "Hello!",
			"log_level": config.LogLevel,
			"host":      config.Host,
		})
	})

	// 启动服务
	addr := ":8080"
	log.Printf("服务启动在 %s", addr)
	log.Println("使用 curl -X POST http://localhost:8080/api/reload 重载配置")
	log.Println("使用 curl http://localhost:8080/api/config 查看当前配置")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
