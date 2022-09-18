var state = false

function toggleModal(Modal, obj){
    state = !state
    if (state) {
        Modal(obj)
        document.documentElement.style.setProperty('--scroll-y', window.scrollY + 'px')
        document.body.style.position = 'fixed'
    } else {
        const body = document.body;
        const scrollY = document.documentElement.style.getPropertyValue('--scroll-y')
        body.style.position = '';
        body.style.top = '';
        window.scrollTo(0, parseInt(scrollY));
    }
    document.querySelector('#modalOverlay').classList.toggle('modal-hidden')
    document.querySelector('#modal').classList.toggle('modal-hidden')
    document.querySelector('body').classList.toggle('modalOpen')
    var e = document.querySelector('#modalBody input, #modalBody select')
    if (e != null) e.focus()
}

function NewInput(parent, type, id, text, change, value){
    var label = document.createElement('label')
    var input = document.createElement('input')
    label.htmlFor = id
    label.innerText = text
    input.type = type
    input.id = id
    input.onchange = change
    input.value = value
    parent.append(label, input)
}

function NewSelect(parent, id, text, change){
    var label = document.createElement('label')
    var select = document.createElement('select')
    label.htmlFor = id
    label.innerText = text
    select.id = id
    select.onchange = change
    parent.append(label, select)
}

function NewButton(parent, text, click){
    var btn = document.createElement('button')
    btn.innerText = text
    btn.onclick = click
    parent.append(btn)
}

function NewDetail(parent, key, value){
    var p = document.createElement('p')
    p.innerText = `${key}: ${value}`
    parent.append(p)
}

document.querySelector('#modal').addEventListener('keydown', (e) => {
    if(e.code == 'Escape') toggleModal()
    else if(e.code == 'Enter') document.querySelector('#modalFooter > button').click()
})