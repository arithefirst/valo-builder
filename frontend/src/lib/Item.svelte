<script lang="js">
    export let weapon
    export let type
    let href

    // Import stores from stores.js
    import { classic } from './stores.js'
    import { shorty } from './stores.js'
    import { frenzy } from './stores.js'
    import { ghost } from './stores.js'
    import { sheriff } from './stores.js'
    import { stinger } from './stores.js'
    import { spectre } from './stores.js'
    import { bucky } from './stores.js'
    import { judge } from './stores.js'
    import { bulldog } from './stores.js'
    import { guardian } from './stores.js'
    import { phantom } from './stores.js'
    import { vandal } from './stores.js'
    import { marshal } from './stores.js'
    import { operator } from './stores.js'
    import { ares } from './stores.js'
    import { odin } from './stores.js'
    import { melee } from './stores.js'
    import { outlaw } from './stores.js'

    const weapons = {
        "classic": classic, "shorty": shorty, "frenzy": frenzy,
        "ghost": ghost, "sheriff": sheriff, "stinger": stinger,
        "spectre": spectre, "bucky": bucky, "judge": judge,
        "bulldog": bulldog, "guardian": guardian, "phantom": phantom,
        "vandal": vandal, "marshal": marshal, "operator": operator,
        "ares": ares, "odin": odin, "melee": melee, "outlaw": outlaw
    }

    weapons[weapon].subscribe((val) => {
        href = val.src
    })
</script>

<div id={weapon}>
    {#if type === "pistol" || type === "smg" }
        <svg width="16.519823788vw" viewBox="0 0 263 109" xmlns="http://www.w3.org/2000/svg">
            <rect width="263" height="109" fill="#FFFFFF1E" />
            <path d="M0,0 L0,109 L263,109 L263,0 M1,1 L1,108 L262,108 L262,1" fill="#FFFFFF32" fill-rule="evenodd"/>
            <path d="M1,1 L1,28 L28,1" fill="#FFFFFF28"/>
            <image {href} width="70%" height="70%" x="15%" y="15%"/>
            <!--Convert weapon name to uppercase before displaying-->
            <text x="3" y="104" font-size="11" fill="white">{weapon.toUpperCase()}</text>
        </svg>
    {/if}
    {#if type === "shotgun" || type === "rifle" || type === "knife" || type === "sniper" || type === "lmg"}
        <svg width="16.519823788vw" viewBox="0 0 263 109" xmlns="http://www.w3.org/2000/svg">
            <rect width="263" height="109" fill="#FFFFFF1E" />
            <path d="M0,0 L0,109 L263,109 L263,0 M1,1 L1,108 L262,108 L262,1" fill="#FFFFFF32" fill-rule="evenodd"/>
            <path d="M1,1 L1,28 L28,1" fill="#FFFFFF28"/>
            <image {href} width="90%" height="90%" x="5%" y="5%"/>
            <!--Convert weapon name to uppercase before displaying-->
            <text x="3" y="104" font-size="11" fill="white">{weapon.toUpperCase()}</text>
        </svg>
    {/if}
</div>

<style>
    svg {
        margin-bottom: 0.82599118942vw;
        font-family: monospace;
        cursor: pointer;
    }
</style>