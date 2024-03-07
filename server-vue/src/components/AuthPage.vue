<template>
    <div id="authPage">
        <input id="username" placeholder="Enter Username">
        <input id="password" type="password" placeholder="Enter Password">
        <button id="login" @click="authUser().then((authStatus) => authenticated = authStatus)">Log In</button>
        <p>Authenticated: {{ authenticated }}</p>
    </div>
</template>

<script>
export default {
    name: 'AuthPage',
    data() {
        return {
            password: '',
            authenticated: false,
        }
    },
    methods: {
        digestMessage: async function(message) {
            const encoder = new TextEncoder()
            const data = encoder.encode(message)
            const hashBuffer = await crypto.subtle.digest("SHA-256", data)
            const hashArray = Array.from(new Uint8Array(hashBuffer))
            const hashHex = hashArray
                .map((b) => b.toString(16).padStart(2, "0"))
                .join("")

            return hashHex
        },

        authUser: async function() {
            var username = document.getElementById('username').value
            var pass = document.getElementById('password').value
            
            const digestHex = await this.digestMessage(pass)

            try {
                const response = await this.$axios.post('/api/auth', {
                    username: username,
                    passwordHash: digestHex,
                })

                console.log(response)
                
                return true
            } catch(err) {
                console.log(err)

                return false
            }
        }
    }
}
</script>

<style>
#authPage {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}

#username {
    position: absolute;
    margin-bottom: 20px;
}

#password {
    margin-top: 40px;
    margin-bottom: 25px;
    position: relative;
}

#login {
    width: 100px;
    height: 30px;
    display: block;
    margin: auto;
}

input {
    width: 200px;
    height: 20px;
    text-align: center;
}
</style>