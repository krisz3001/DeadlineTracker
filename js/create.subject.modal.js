function SubjectModalBuilder(){
    ClearModal()
    SetModalTitle('New Subject')
    NewInput('text', 'subjectName', 'Subject Name', inputChanged)
    NewButton('Add', addNewSubject)
}

function SubjectModal(){
    SubjectModalBuilder()
    SetInputs()
}

function addNewSubject(){
    if(CheckForm()) return
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:3556/subjects", true)
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
        subjectname: inputs[0].value
    }))
}