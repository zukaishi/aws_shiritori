var app = new Vue({ 
    el: '#app',
    data: {
        message: 'Hello Vue!',
        errors: [],
        name1: null,
        name2: null,
        result: null
    },
    methods:{
      checkForm: function (e) {
        if (this.name1 && this.name2) {
          return true;
        }
        this.errors = [];
  
        if (!this.name1) {
          this.errors.push('Name1 required.');
        }
        if (!this.name2) {
          this.errors.push('Name2 required.');
        }
  
        e.preventDefault();
      }
    },
    mounted () {
      axios
        .get('https://api.coindesk.com/v1/bpi/currentprice.json')
        .then(response => (this.result = response))
    }
});
