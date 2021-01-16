var app = new Vue({ 
    el: '#app',
    data: {
        errors: [],
        name1: null,
        name2: null,
        result: []
    },
    methods:{
      checkForm: function (e) {
        e.preventDefault();
        this.errors = [];
        let results = []

        /*
        if (!this.name1) {
          this.errors.push('Name1 required.');
        }
        if (!this.name2) {
          this.errors.push('Name2 required.');
        }
*/
        axios.get('https://jiehmlfyck.execute-api.ap-northeast-1.amazonaws.com/default/api-test')
          .then((response) => {
            str = response.data.body.split(",");
            str.forEach(function(str){
              results.push(str);
            })
          })
          this.result = results
        }
    },
    mounted () {
    }
});
