var application = new Vue({
    el : "#crudApp",
    data :{
        temp:"alex",
        object: [],
    },
    delimiters: ['[[', ']]'],
    methods: {
        fetchObject : function(){
            axios.post('/fetchobj').then (function(response){
                application.object = response.data;
            });
        },
        getObjsize: function(){
            var len = 0;
            if (this.object.length%2 == 1){
                len = (this.object.length-1)/2;
            }
            else{
                len =  this.object.length/2;
            }
            return len;
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