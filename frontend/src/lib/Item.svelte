<script>
    export let weapon
    export let url
</script>

<svg id="" width="16.519823788vw" viewBox="0 0 263 109" xmlns="http://www.w3.org/2000/svg">
    <rect width="263" height="109" fill="#245269" />
    <path d="M0,0 L0,109 L263,109 L263,0 M1,1 L1,108 L262,108 L262,1" fill="#386b8b" fill-rule="evenodd"/>
    <path d="M1,1 L1,28 L28,1" fill="#3b5e76"/>
    <text class="text" x="3" y="104" font-size="11" fill="white">{weapon.toUpperCase()}</text>
    <style>
        .text {
            font-family: cursive;
        }
    </style>
</svg>

<style>
    svg {
        margin-bottom: 0.82599118942vw;
        font-family: cursive;
    }
</style>