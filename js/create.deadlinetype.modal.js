function DeadlineTypeModalBuilder(){
    ClearModal()
    SetModalTitle('New Deadline Type')
    NewInput('text', 'typeInput', 'Deadline Type', inputChanged)
    NewButton('Add', addDeadlineType)
}

function DeadlineTypeModal(){
    DeadlineTypeModalBuilder()
    SetInputs()
}

function addDeadlineType(){
    if(CheckForm()) return
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:3556/deadlinetypes", true)
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
        deadlinetypename: inputs[0].value
    }))
}