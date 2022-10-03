function DetailsModalBuilder(i){
    ClearModal()
    SetModalTitle('Details')
    NewInput('datetime-local', 'deadlineInput', 'Deadline', inputChanged, Deadlines.data[i].deadline)
    NewSelect('deadlineTypeSelect', 'Type', inputChanged)
    NewSelect('subjectSelect', 'Subject', inputChanged)
    NewInput('text', 'topicInput', 'Topic', inputChanged, Deadlines.data[i].topic)
    NewInput('text', 'commentsInput', 'Comments', inputChanged, Deadlines.data[i].comments)
    NewInput('checkbox', 'isFixed', 'Fixed', null, Deadlines.data[i].fixed)
    NewButton('Update', function() {updateDeadline(i)})
    NewButton('Delete', function() {deleteDeadline(i)})
}

function DetailsModal(obj){
    var index
    if (obj != null) index = Array.from(document.querySelectorAll('.top, .later')).indexOf(obj)
    DetailsModalBuilder(index)
    SetInputs()
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
    if(CheckForm()) return
    var xhr = new XMLHttpRequest()
    xhr.open("PATCH", `/deadlines/${id}`, true)
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
        comments: inputs[4].value,
        fixed: inputs[5].checked*1
    }))
}

function deleteDeadline(i){
    if(!confirm('Are you sure?')) return
    var id = Deadlines.data[i].id
    var xhr = new XMLHttpRequest()
    xhr.open("DELETE", `/deadlines/${id}`, true)
    xhr.onreadystatechange = () => {
        if(xhr.readyState != 4) return
        if(xhr.status == 200){
            Reload()
            toggleModal()
        }
    }
    xhr.send()
}