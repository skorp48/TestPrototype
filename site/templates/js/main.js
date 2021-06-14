var application = new Vue({
    el : "#crudApp",
    data :{
        temp:"alex",
        object: [] 
    },
    delimiters: ['[[', ']]'],
    methods: {
        fetchObject : function(){
            axios.post('/fetchobj').then (function(response){
                application.object = response.data;
            });
        }
    },
    created: function () {
        this.fetchObject();
    }
})

    // параллакс
    let bg = document.querySelector('.Header-image');
	window.addEventListener('mousemove', function(e) {
	let x = e.clientX / window.innerWidth;
	let y = e.clientY / window.innerHeight;  
	bg.style.transform = 'translate(-' + x * 50 + 'px, -' + y * 50 + 'px)';});