var list = [{"twitchName":"","fileName":"","alternativeName":null}];

var app = new Vue({
    el: '#app',
    data: {
        search: '',
        games: list,
        edit: {
            index: 0,
            input: "",
            show: false,
        }
    },
    computed: {
        filtered: function () {
            return this.games.filter(game => {
                return game.twitchName.toLowerCase().indexOf(this.search.toString().toLowerCase()) > -1
            })
        },
    },
    methods: {
        openEdit: function(name){
            for (let i = 0; i < this.games.length; i++) {
                const game = this.games[i];
                if(game.twitchName == name){
                    this.edit.index = i
                    this.edit.show = true
                    return
                }
            }
        },
        addAlt: function () {
            if (this.games[this.edit.index].alternativeName) {
                this.games[this.edit.index].alternativeName.push(this.edit.input);
            } else {
                this.games[this.edit.index].alternativeName = [this.edit.input];
            }
            this.edit.input = ""
        },
        removeAlt: function (index) {
            this.games[this.edit.index].alternativeName.splice(index, 1);
        },
        updateGame: function () {
            console.log(this.games[this.edit.index])
            let url = "http://localhost:8000/gamelist"
            fetch(url, {
                    body: JSON.stringify(this.games[this.edit.index]), // must match 'Content-Type' header
                    cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
                    headers: {
                    'content-type': 'application/json'
                    },
                    method: 'POST', // *GET, POST, PUT, DELETE, etc.
                })
                .then(response => {})
        },
    },
    created () {
        fetch("http://localhost:8000/twitch/gamelist")
            .then(response => response.json())
            .then(json => {
                this.games = json
            })
    }
});