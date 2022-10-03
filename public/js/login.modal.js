function LoginModalBuilder(){
    ClearModal()
    SetModalTitle('Login')
    NewInput('text', 'username', 'Username', inputChanged)
    NewInput('password', 'password', 'Password', inputChanged)
    NewButton('Login', Login)
}

function LoginModal(){
    LoginModalBuilder()
    SetInputs()
}

function Login(){
    if(CheckForm()) return
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "/login", true)
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.onreadystatechange = () => {
        if(xhr.readyState != 4) return
        if(xhr.status == 200){
            location.href = ""
        }
        if(xhr.status == 401){
            NewError('incorrect username or password')
        }
    }
    xhr.send(JSON.stringify({
        username: inputs[0].value,
        password: inputs[1].value
    }))
}