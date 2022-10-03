var state = false
var modal = document.querySelector('#modalBody')
var modalTitle = document.querySelector('#modalTitle')
var modalFooter = document.querySelector('#modalFooter')
var inputs

function toggleModal(Modal, obj){
    state = !state
    if (state) {
        Modal(obj)
        document.querySelector('.dropdown-content').style.display = 'none'
        document.documentElement.style.setProperty('--scroll-y', window.scrollY + 'px')
        document.body.style.position = 'fixed'
    } else {
        document.querySelector('.dropdown-content').style.display = ''
        const body = document.body;
        const scrollY = document.documentElement.style.getPropertyValue('--scroll-y')
        body.style.position = '';
        body.style.top = '';
        window.scrollTo(0, parseInt(scrollY));
    }
    document.querySelector('#modalOverlay').classList.toggle('modal-hidden')
    document.querySelector('#modal').classList.toggle('modal-hidden')
    document.querySelector('body').classList.toggle('modalOpen')
    inputs[0].focus()
}

function ClearModal(){
    modalFooter.innerHTML = ''
    modal.innerHTML = ''
}

function SetModalTitle(title){
    modalTitle.innerText = title
}

function SetInputs(){
    inputs = document.querySelectorAll('#modal input, #modal select, #modalFooter input')
}

function NewInput(type, id, text, change, value){
    var label = document.createElement('label')
    var input = document.createElement('input')
    label.htmlFor = id
    label.innerText = text
    input.type = type
    input.id = id
    if(type == 'checkbox'){
        var div = document.createElement('div')
        div.style.position = 'absolute'
        div.style.left = '10px'
        label.style.display = 'inline-block'
        label.style.marginRight = '10px'
        value == 1 ? input.checked = true : input.checked = false
        div.append(label, input)
        modalFooter.append(div)
    } else{
        if(value != undefined) input.value = value
        input.oninput = change
        modal.append(label, input)
    }
}

function NewSelect(id, text, change){
    var label = document.createElement('label')
    var select = document.createElement('select')
    label.htmlFor = id
    label.innerText = text
    select.id = id
    select.oninput = change
    modal.append(label, select)
}

function NewButton(text, click){
    var btn = document.createElement('button')
    btn.innerText = text
    btn.onclick = click
    modalFooter.append(btn)
}

function NewDetail(key, value){
    var p = document.createElement('p')
    p.innerText = `${key}: ${value}`
    modal.append(p)
}

function NewError(text) {
    var p = document.querySelector('.errorMessage')
    if(p == undefined) p = document.createElement('p')
    p.innerText = text[0].toUpperCase() + text.substring(1, text.length)
    p.classList.add('errorMessage')
    modal.append(p)
}

document.querySelector('#modal').addEventListener('keydown', (e) => {
    if(e.code == 'Escape') toggleModal()
    else if(e.code == 'Enter') document.querySelector('#modalFooter > button').click()
})

function inputChanged(){
    this.classList.remove('wrongInput')
    document.querySelector(`label[for=${this.id}]`).classList.remove('wrongInputLabel')
}

function CheckForm(){
    var wrong = false
    for (let i = 0; i < inputs.length; i++) {
        if(inputs[i].value == '0' || inputs[i].value == '' &&inputs[i].type != 'checkbox'){
            inputs[i].classList.add('wrongInput')
            document.querySelector(`label[for="${inputs[i].id}"]`).classList.add('wrongInputLabel')
            wrong = true  
        }
    }
    return wrong
}