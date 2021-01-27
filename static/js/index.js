var app = new Vue({ 
    el: '#app',
    data: {
        errors: [],
        result: [],
        name1: null,
        name2: null
    },
    methods:{
      request: function (e) {
        e.preventDefault();
        this.errors = [];
        let results = []

        if (!this.name1) {
          this.errors.push('Name1 required.');
        }
        let url =""
        if (!this.name2) {
          url = "https://2wb8kl0nf6.execute-api.ap-northeast-1.amazonaws.com/default/comprised" + "?name1=" + this.name1 
          this.errors.push('Name2 required.');
        } else {
          url = "https://jiehmlfyck.execute-api.ap-northeast-1.amazonaws.com/default/api-test" + "?name1=" + this.name1 + "&name2=" + this.name2 
        }

        axios.get(url)
          .then((response) => {
            str = response.data.body.split(",");
            str.forEach(function(str){
              results.push(str);
            })
          })
          this.result = results
        },
        clear: function() {
          this.name1 = ""
          this.name2 = ""
        }
    },
    mounted () {
    }
});
