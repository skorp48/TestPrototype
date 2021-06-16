var application = new Vue({
    el : "#TestResult",
    data :{
        object: [],
        clicked: [false,false,false]
    },
    delimiters: ['[[', ']]'],
    methods: {
        clickFu:function(num){
            this.$set(this.clicked, num, !this.clicked[num])
        }

        // fetchObject : function(){
        //     axios.post('/fetchobj').then (function(response){
        //         application.object = response.data;
        //     });
        // }
    }
    // created: function () {
    //     this.fetchObject();
    // }
})

    // параллакс
    let bg = document.querySelector('.Header-image');
	window.addEventListener('mousemove', function(e) {
	let x = e.clientX / window.innerWidth;
	let y = e.clientY / window.innerHeight;  
	bg.style.transform = 'translate(-' + x * 50 + 'px, -' + y * 50 + 'px)';});