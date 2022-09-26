function SignupModalBuilder(){
    ClearModal()
    SetModalTitle('Sign Up')
    NewInput('text', 'username', 'Username', inputChanged)
    NewInput('password', 'password', 'Password', inputChanged)
    NewInput('password', 'secret', 'Secret', inputChanged)
    NewButton('Sign up', SignUp)
}

function SignupModal(){
    SignupModalBuilder()
    SetInputs()
}

function SignUp(){
    if(CheckForm()) return
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "/signup", true)
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.onreadystatechange = () => {
        if(xhr.readyState != 4) return
        if(xhr.status == 200){
            location.href = ""
			toggleModal()
        }
      if(xhr.status == 409){
        NewError(xhr.response)
      }
      if(xhr.status == 401){
        NewError(xhr.response)
      }
    }
    xhr.send(JSON.stringify({
        username: inputs[0].value,
        password: inputs[1].value,
        secret: inputs[2].value
    }))
}