function DeadlineModalBuilder(){
    document.querySelector('#modalTitle').innerText = 'New Deadline'
    var modal = document.querySelector('#modalBody')
    var modalFooter = document.querySelector('#modalFooter')
    modalFooter.innerHTML = ''
    modal.innerHTML = ''
    NewInput(modal, 'datetime-local', 'deadlineDate', 'Choose date', selectChanged)
    NewSelect(modal, 'deadlineTypeSelect', 'Type', selectChanged)
    NewSelect(modal, 'subjectSelect', 'Subject', selectChanged)
    NewInput(modal, 'text', 'topicInput', 'Topic')
    NewInput(modal, 'text', 'commentsInput', 'Comments')
    NewButton(modalFooter, 'Add', addNewDeadline)
}

function DeadlineModal(){
    DeadlineModalBuilder()
    var subjectSelect = document.getElementById('subjectSelect')
    var deadlineTypeSelect = document.getElementById('deadlineTypeSelect')
    var inputs = document.querySelectorAll('#deadlineDate, #topicInput, #commentsInput')
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
    var inputs = document.querySelectorAll('#modal input, #modal select')
    var wrong = false
    for (let i = 0; i < 3; i++) {
        if(inputs[i].value == '0' || inputs[i].value == ''){
            inputs[i].classList.add('wrongInput')
            document.querySelector(`label[for=${inputs[i].id}]`).classList.add('wrongInputLabel')
            wrong = true  
        }
    }
    if(wrong) return
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

function selectChanged(){
    this.classList.remove('wrongInput')
    document.querySelector(`label[for=${this.id}]`).classList.remove('wrongInputLabel')
}