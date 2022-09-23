function DeadlineModalBuilder(){
    ClearModal()
    SetModalTitle('New Deadline')
    NewInput('datetime-local', 'deadlineDate', 'Choose date', inputChanged)
    NewSelect('deadlineTypeSelect', 'Type', inputChanged)
    NewSelect('subjectSelect', 'Subject', inputChanged)
    NewInput('text', 'topicInput', 'Topic', inputChanged)
    NewInput('text', 'commentsInput', 'Comments', inputChanged)
    NewButton('Add', addNewDeadline)
}

function DeadlineModal(){
    DeadlineModalBuilder()
    SetInputs()
    var subjectSelect = document.getElementById('subjectSelect')
    var deadlineTypeSelect = document.getElementById('deadlineTypeSelect')
    inputs.forEach(e => {
        e.value = ""
    });
    subjectSelect.innerHTML = `<option value="0">-- choose --</option>`
    deadlineTypeSelect.innerHTML = `<option value="0">-- choose --</option>`
    var elements = document.querySelectorAll('.modal input, .modal select, .modal label')
    elements.forEach(e => {
        e.classList.remove('wrongInput')
        e.classList.remove('wrongInputLabel')
    })
    for (let i = 0; i < SubjectsCache.data.length; i++) {
        subjectSelect.innerHTML += `<option value="${SubjectsCache.data[i].subjectkey}">${SubjectsCache.data[i].subjectname}</option>`
    }
    for (let i = 0; i < TypesCache.data.length; i++) {
        deadlineTypeSelect.innerHTML += `<option value="${TypesCache.data[i].deadlinetypeid}">${TypesCache.data[i].deadlinetypename}</option>`
    }
}

function addNewDeadline(){
    if(CheckForm()) return
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:3556/deadlines", true)
    xhr.setRequestHeader('Content-Type', 'application/json')
    xhr.onreadystatechange = () => {
        if(xhr.readyState != 4) return
        if(xhr.status == 200){
            Reload()
            toggleModal()
        }
    }
    xhr.send(JSON.stringify({
        typeid: inputs[1].value*1,
        deadline: inputs[0].value.replace('T',' '),
        subjectid: inputs[2].value*1,
        topic: inputs[3].value,
        comments: inputs[4].value
    }))
}