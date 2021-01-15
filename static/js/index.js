var app = new Vue({ 
    el: '#app',
    data: {
        errors: [],
        name1: null,
        name2: null,
        result: null
    },
    methods:{
      checkForm: function (e) {
        e.preventDefault();
        this.errors = [];
/*
        if (!this.name1) {
          this.errors.push('Name1 required.');
        }
        if (!this.name2) {
          this.errors.push('Name2 required.');
        }
*/


        }
    },
    mounted () {
      axios.get('https://jiehmlfyck.execute-api.ap-northeast-1.amazonaws.com/default/api-test')
        .then((response) => {
          this.result = response.data.body;
        })
    }
});
