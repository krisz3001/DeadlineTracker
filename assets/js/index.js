var DeadlinesCache
var SubjectsCache
var TypesCache

function NewSoonItem(deadline, type, subject, topic, comments, fixed){
    var div = document.createElement('div')
    div.classList.add('top')
    if(fixed == 0) div.classList.add('notfixed')
    div.onclick = function() {toggleModal(DetailsModal, this)}
    var dl = document.createElement('p')
    dl.classList.add('deadline')
    dl.innerText = deadline
    var hrTop = document.createElement('hr')
    var hrBottom = document.createElement('hr')
    var typeText = document.createElement('p')
    typeText.classList.add('type')
    typeText.innerText = type
    var subjectText = document.createElement('p')
    subjectText.classList.add('subject')
    subjectText.innerText = subject
    var topicText = document.createElement('p')
    topicText.classList.add('topic')
    topicText.innerText = topic
    var commentsText = document.createElement('p')
    commentsText.classList.add('comments')
    commentsText.innerText = comments
    div.append(dl, hrTop, typeText, subjectText, topicText, hrBottom, commentsText)
    document.querySelector('#soon').append(div)
    //return `<div class="top"><p class="deadline">${deadline}</p><hr><p class="type">${type}</p><p class="subject">${subject}</p><p class="topic">${topic}</p><hr><p class="comments">${comments}</p></div>`
}

function NewLaterItem(deadline, type, subject, fixed){
    var div = document.createElement('div')
    div.classList.add('later')
    if(fixed == 0) div.classList.add('notfixed')
    div.onclick = function() {toggleModal(DetailsModal, this)}
    var dl = document.createElement('p')
    dl.classList.add('deadlineLater', 'left')
    dl.innerText = deadline
    var typeText = document.createElement('p')
    typeText.classList.add('type')
    typeText.innerText = type
    var subjectText = document.createElement('p')
    subjectText.classList.add('subject', 'right')
    subjectText.innerText = subject
    div.append(dl, typeText, subjectText)
    document.querySelector('#deadlinesLater').append(div)
    //return `<div class="later" onclick="function() {toggleModal(DetailsModal, this)}"><p class="deadlineLater left">${deadline}</p><p class="type">${type}</p><p class="subject right">${subject}</p></div>`
}
function Reload(){
    var soon = document.getElementById('soon')
    var later = document.getElementById('deadlinesLater')
    fetch("/deadlines")
    .then(res=> res.json())
    .then(r=>{
        soon.innerHTML = ""
        later.innerHTML = ""
        Deadlines = r
        for (let i = 0; i < r.data.length; i++) {
            if(i < 3) NewSoonItem(ToPrettyDate(r.data[i].deadline), r.data[i].type, r.data[i].subject, r.data[i].topic, r.data[i].comments, r.data[i].fixed)
            else NewLaterItem(ToPrettyDate(r.data[i].deadline), r.data[i].type, r.data[i].subject, r.data[i].fixed)
        }
    })
    fetch('/subjects')
    .then(res => res.json())
    .then(r => {
        SubjectsCache = r
    })
    fetch('/deadlinetypes')
    .then(res => res.json())
    .then(r => {
        TypesCache = r
    })
}

function ToPrettyDate(date){
    var options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
    return new Date(date).toLocaleDateString('hu-HU', options)
}

function Logout(){
    fetch('/logout')
    .then(res => res.text())
    .then(() => {
        location.href = ""
        document.cookie = "Token=; Max-Age=-9999999"
    })
}

Reload()