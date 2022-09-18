function DetailsModalBuilder(i){
    document.querySelector('#modalTitle').innerText = 'Details'
    var modal = document.querySelector('#modalBody')
    var modalFooter = document.querySelector('#modalFooter')
    modalFooter.innerHTML = ''
    modal.innerHTML = ''
    NewInput(modal, 'datetime-local', 'deadlineInput', 'Deadline', null, Deadlines.data[i].deadline)
    NewSelect(modal, 'deadlineTypeSelect', 'Type', selectChanged)
    NewSelect(modal, 'subjectSelect', 'Subject', selectChanged)
    NewInput(modal, 'text', 'topicInput', 'Topic', null, Deadlines.data[i].topic)
    NewInput(modal, 'text', 'commentsInput', 'Comments', null, Deadlines.data[i].comments)
    NewButton(modalFooter, 'Update', function() {updateDeadline(i)})
    NewButton(modalFooter, 'Delete', function() {deleteDeadline(i)})
}

function DetailsModal(obj){
    var index
    if (obj != null) index = Array.from(document.querySelectorAll('.top, .later')).indexOf(obj)
    DetailsModalBuilder(index)
    var typeIndex
    var subjectIndex
    var deadlineTypeSelect = document.getElementById('deadlineTypeSelect')
    var subjectSelect = document.querySelector('#subjectSelect')
    for (let i = 0; i < TypesCache.data.length; i++) {
        if(TypesCache.data[i].deadlinetypename == Deadlines.data[index].type) typeIndex = i
        deadlineTypeSelect.innerHTML += `<option value="${TypesCache.data[i].deadlinetypeid}">${TypesCache.data[i].deadlinetypename}</option>`
    }
    deadlineTypeSelect.selectedIndex = typeIndex
    for (let i = 0; i < SubjectsCache.data.length; i++) {
        if(SubjectsCache.data[i].subjectname == Deadlines.data[index].subject) subjectIndex = i
        subjectSelect.innerHTML += `<option value="${SubjectsCache.data[i].subjectkey}">${SubjectsCache.data[i].subjectname}</option>`
    }
    subjectSelect.selectedIndex = subjectIndex
}

function updateDeadline(i){
    var id = Deadlines.data[i].id
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
    xhr.open("PATCH", `http://localhost:3556/deadlines/${id}`, true)
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

function deleteDeadline(i){
    if(!confirm('Are you sure?')) return
    var id = Deadlines.data[i].id
    var xhr = new XMLHttpRequest()
    xhr.open("DELETE", `http://localhost:3556/deadlines/${id}`, true)
    xhr.onreadystatechange = () => {
        if(xhr.readyState != 4) return
        if(xhr.status == 200){
            Reload()
            toggleModal()
        }
    }
    xhr.send()
}