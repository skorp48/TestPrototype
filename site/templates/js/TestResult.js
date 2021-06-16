var application = new Vue({
    el : "#TestResult",
    data :{
        result: [],
        clicked: [false,false,false]
    },
    delimiters: ['[[', ']]'],
    methods: {
        clickFu:function(num){
            this.$set(this.clicked, num, !this.clicked[num])
        },
        fetchResult : function(){
            axios.post('/fetchres').then (function(response){
                application.result = response.data;
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