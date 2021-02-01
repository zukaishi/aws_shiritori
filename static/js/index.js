var app = new Vue({ 
    el: '#app',
    data: {
        errors: [],
        result: [],
        name1: null,
        name2: null,
        mode:0 // 0 しりとり,1 どこかに文字列,2 最初に文字列,3 最後に文字列
    },
    methods:{
      request: function (e) {
        e.preventDefault();
        this.errors = [];
        let name1 = this.name1
        let name2 = this.name2

        if (!this.name1) {
          this.errors.push('Name1 required.');
        } else {
          this.name1 = this.hiraToKana(this.name1)
        }
        if(this.name2) {
          this.name2 = this.hiraToKana(this.name2)
        }

        let url =""
        if (!this.name2) {
          url = "https://2wb8kl0nf6.execute-api.ap-northeast-1.amazonaws.com/default/comprised" + "?name1=" + name1
        } else {
          url = "https://jiehmlfyck.execute-api.ap-northeast-1.amazonaws.com/default/api-test" + "?name1=" + name1+ "&name2=" + name2
        }

        let mode = this.mode 
        let results = []
        axios.get(url)
          .then((response) => {
            let body = response.data.body.split(",");
            // https://ja.stackoverflow.com/questions/72150/vue%E3%81%AEforeach%E3%81%AE%E4%B8%AD%E3%81%A7%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%81%8C%E5%91%BC%E3%81%B9%E3%81%AA%E3%81%84
            body.forEach((value) => {
              if( value ) {

                let [no,name] = value.split(':');
                if ( mode == 2 ) {
                  if( name1 !== name.slice(0,name1.length) ) {
                    name = ""
                  }
                } else if ( mode == 3 ) {
                  if( name1 !== name.slice(name1.length*-1) ) {
                    name = ""
                  }
                }
                if (name ) {
                  results.push(  name + "(" + this.kanaToHira(name) + ")");
                }
              }
            });
          })
        this.result = results
      },
      clear: function() {
        this.name1 = ""
        this.name2 = ""
      },
      kanaToHira: function(str) {
        return str.replace(/[\u30a1-\u30f6]/g, function(match) {
            var chr = match.charCodeAt(0) - 0x60;
            return String.fromCharCode(chr);
        });
      },
      hiraToKana: function(str) {
        return str.replace(/[\u3041-\u3096]/g, function(match) {
            var chr = match.charCodeAt(0) + 0x60;
            return String.fromCharCode(chr);
        });
      }
    }
});
