<script>
    export let weapon
    export let href
    export let type

    if (href === "") {
        const defaults = new Map()
        defaults.set('classic', "https://media.valorant-api.com/weaponskinchromas/4b2d5b4f-4955-4208-286c-abadec250cdd/fullrender.png")
        defaults.set('shorty', "https://media.valorant-api.com/weaponskinchromas/95608504-4c8b-1408-1612-0f8200421c49/fullrender.png")
        defaults.set('frenzy', "https://media.valorant-api.com/weaponskinchromas/dc99ed5a-4d75-87a0-c921-75963ea3c1e1/fullrender.png")
        defaults.set('ghost', "https://media.valorant-api.com/weaponskinchromas/947a28b6-4e0f-61fb-e795-bc9a5e7b7129/fullrender.png")
        defaults.set('sheriff', "https://media.valorant-api.com/weaponskinchromas/5a59bd61-48a9-af61-c00f-4aa21deca9a8/fullrender.png")
        defaults.set('stinger', "https://media.valorant-api.com/weaponskinchromas/31bb2115-4c62-d37c-43c4-11b8fee7f212/fullrender.png")
        defaults.set('spectre', "https://media.valorant-api.com/weaponskinchromas/a9aaccca-4cdc-02ea-1d7e-89bbacecc0e2/fullrender.png")
        defaults.set('bucky', "https://media.valorant-api.com/weaponskinchromas/3d8ffcfe-4786-0180-42d7-e1be18dd1cab/fullrender.png")
        defaults.set('judge', "https://media.valorant-api.com/weaponskinchromas/b71ae8d6-44bb-aa4c-0d2a-dc9ed9e66410/fullrender.png")
        defaults.set('bulldog', "https://media.valorant-api.com/weaponskinchromas/bf35f404-4a14-6953-ced2-5bafd21639a0/fullrender.png")
        defaults.set('guardian', "https://media.valorant-api.com/weaponskinchromas/0f934388-418a-a9e7-42a7-21b27402e46c/fullrender.png")
        defaults.set("phantom", "https://media.valorant-api.com/weaponskinchromas/52221ba2-4e4c-ec76-8c81-3483506d5242/fullrender.png")
        defaults.set('vandal', "https://media.valorant-api.com/weaponskinchromas/19629ae1-4996-ae98-7742-24a240d41f99/fullrender.png")
        defaults.set('melee', "https://media.valorant-api.com/weaponskinchromas/cac83e5c-47a1-3519-5420-1db1fdbc4892/fullrender.png")
        defaults.set('marshal', "https://media.valorant-api.com/weaponskinchromas/1afec971-4170-f29b-1c94-07a0eff270ab/fullrender.png")
        defaults.set('outlaw', "https://media.valorant-api.com/weaponskinchromas/66c8d241-4f7c-6652-3aaa-51bafffbd493/fullrender.png")
        defaults.set('operator', "https://media.valorant-api.com/weaponskinchromas/4914f50d-49f9-6424-ca80-9486c45a138d/fullrender.png")
        defaults.set('ares', "https://media.valorant-api.com/weaponskinchromas/b33de820-4061-8b85-31ce-808f1a2c58f5/fullrender.png")
        defaults.set('odin', "https://media.valorant-api.com/weaponskinchromas/2f93861d-4b2f-2175-af0c-3ba0c736e257/fullrender.png")
        href = defaults.get(weapon)
    }
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
        font-family: cursive;
        cursor: pointer;
    }
</style>