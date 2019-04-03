console.log('BEGIN');

window.onload = function() {
    console.log('window.onload');

    var createElement = function(tag, id = '', cls = null) {
        var newElement = document.createElement(tag);
        newElement.id = id;
        if(cls)
            newElement.className = cls;
        return newElement;
    }

    document.body.appendChild(createElement('div', 'main', 'mainClass'));

    var list = createElement('ul', 'list');
    for(var i = 0; i < 3; i++) {
        var li = document.createElement('li');
        li.name = 'li' + i;
        li.innerHTML = 'Mast Know ' + i;
        list.appendChild(li);
    }

    var mainDiv = document.getElementById('main');
    mainDiv.appendChild(list);    
}