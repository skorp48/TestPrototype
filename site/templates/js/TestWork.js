var application = new Vue({
    el : "#TestWork",
    data :{
        temp:"alex2",
    },
    delimiters: ['[[', ']]'],
})

    // Параллакс
    let bg = document.querySelector('.Header-image');
	window.addEventListener('mousemove', function(e) {
	let x = e.clientX / window.innerWidth;
	let y = e.clientY / window.innerHeight;  
	bg.style.transform = 'translate(-' + x * 50 + 'px, -' + y * 50 + 'px)';});