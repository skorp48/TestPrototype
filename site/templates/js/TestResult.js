var application = new Vue({
    el : "#TestResult",
    data :{
        result: [],
        clicked: []
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
        },
        ClickCount : function(){
            this.result.length
            for (let index = 1; index < this.result.length; index++) {
                clicked.push(false)
            }
        }
    },
    created: function () {
        this.fetchResult();
    }
})

    // параллакс
    let bg = document.querySelector('.Header-image');
	window.addEventListener('mousemove', function(e) {
	let x = e.clientX / window.innerWidth;
	let y = e.clientY / window.innerHeight;  
	bg.style.transform = 'translate(-' + x * 50 + 'px, -' + y * 50 + 'px)';});