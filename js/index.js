var DeadlinesCache
var SubjectsCache
var TypesCache

function NewSoonItem(deadline, type, subject, topic, comments){
    var div = document.createElement('div')
    div.classList.add('top')
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

function NewLaterItem(deadline, type, subject){
    var div = document.createElement('div')
    div.classList.add('later')
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
    fetch("http://localhost:3556/deadlines")
    .then(res=> res.json())
    .then(r=>{
        soon.innerHTML = ""
        later.innerHTML = ""
        Deadlines = r
        for (let i = 0; i < r.data.length; i++) {
            if(i < 3) NewSoonItem(ToPrettyDate(r.data[i].deadline), r.data[i].type, r.data[i].subject, r.data[i].topic, r.data[i].comments)
            else NewLaterItem(ToPrettyDate(r.data[i].deadline), r.data[i].type, r.data[i].subject)
        }
        var tops = document.querySelectorAll('.top')
        tops[1].style.marginTop = '30px'
        tops[2].style.marginTop = '60px'
    })
    fetch('http://localhost:3556/subjects')
    .then(res => res.json())
    .then(r => {
        SubjectsCache = r
    })
    fetch('http://localhost:3556/deadlinetypes')
    .then(res => res.json())
    .then(r => {
        TypesCache = r
    })
}

function ToPrettyDate(date){
    var options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
    return new Date(date).toLocaleDateString('hu-HU', options)
}

Reload()