// 获取页面上的元素
let loginForm = document.querySelector("form")
let nameInp = document.querySelector(".username")
let pwdInp = document.querySelector(".password")
let errBox = document.querySelector(".error")

// 绑定事件
loginForm.onsubmit = function (e) {
    e.preventDefault()
    // 获取表单信息
    let name = nameInp.value
    let pwd = pwdInp.value

    // 验证用户名和密码
    if (!name || !pwd) {
        return alert("用户名和密码不能为空")
    }

    // 发送 ajax 请求
    let xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:8080/api/v1/login", true)
    xhr.onload = function () {
        console.log(xhr.responseText)
        // 根据响应体中的信息，更新页面
        // let resp = JSON.parse(xhr.responseText)
        // if (resp.code === 1) {
        //     // 登录成功后，跳转到其他页面
        //     window.location.href = "./order.html"
        // } else {
        //     errBox.style.display = 'block'
        // }
    }
    xhr.setRequestHeader("content-type", "application/x-www-form-urlencoded")
    xhr.send('username=' + name + '&password=' + pwd)
}