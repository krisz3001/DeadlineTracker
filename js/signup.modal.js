function SignupModalBuilder(){
    ClearModal()
    SetModalTitle('Sign Up')
    NewInput('text', 'username', 'Username', inputChanged)
    NewInput('password', 'password', 'Password', inputChanged)
    NewButton('Sign up', SignUp)
}

function SignupModal(){
    SignupModalBuilder()
    SetInputs()
}

function SignUp(){
    if(CheckForm()) return
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:3556/signup", true)
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.onreadystatechange = () => {
        if(xhr.readyState != 4) return
        if(xhr.status == 200){
            Reload()
			toggleModal()
        }
      if(xhr.status == 409){
            NewError(xhr.response)
        }
    }
    xhr.send(JSON.stringify({
        username: inputs[0].value,
        password: inputs[1].value
    }))
}