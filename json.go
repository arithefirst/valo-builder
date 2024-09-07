package main

func giveJson() []byte {
	json := []byte(`{
      "uuid": "63e6c2b6-4a8e-869c-3d4c-e38355226584",
      "displayName": "Odin",
      "category": "EEquippableCategory::Heavy",
      "defaultSkinUuid": "f454efd1-49cb-372f-7096-d394df615308",
      "displayIcon": "https://media.valorant-api.com/weapons/63e6c2b6-4a8e-869c-3d4c-e38355226584/displayicon.png",
      "killStreamIcon": "https://media.valorant-api.com/weapons/63e6c2b6-4a8e-869c-3d4c-e38355226584/killstreamicon.png",
      "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HMGPrimaryAsset",
      "weaponStats": {
        "fireRate": 12,
        "magazineSize": 100,
        "runSpeedMultiplier": 0.76,
        "equipTimeSeconds": 1.25,
        "reloadTimeSeconds": 5,
        "firstBulletAccuracy": 0.8,
        "shotgunPelletCount": 1,
        "wallPenetration": "EWallPenetrationDisplayType::High",
        "feature": "EWeaponStatsFeature::ROFIncrease",
        "fireMode": null,
        "altFireType": "EWeaponAltFireDisplayType::ADS",
        "adsStats": { "zoomMultiplier": 1.15, "fireRate": 15.6, "runSpeedMultiplier": 0.76, "burstCount": 1, "firstBulletAccuracy": 0.79 },
        "altShotgunStats": null,
        "airBurstStats": null,
        "damageRanges": [
          { "rangeStartMeters": 0, "rangeEndMeters": 30, "headDamage": 95, "bodyDamage": 38, "legDamage": 32.3 },
          { "rangeStartMeters": 30, "rangeEndMeters": 50, "headDamage": 77.5, "bodyDamage": 31, "legDamage": 26.35 }
        ]
      },
      "shopData": {
        "cost": 3200,
        "category": "Heavy Weapons",
        "shopOrderPriority": 0,
        "categoryText": "Heavy Weapons",
        "gridPosition": { "row": 2, "column": 2 },
        "canBeTrashed": true,
        "image": null,
        "newImage": "https://media.valorant-api.com/weapons/63e6c2b6-4a8e-869c-3d4c-e38355226584/shop/newimage.png",
        "newImage2": null,
        "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HeavyMachineGunPurchase"
      },
      "skins": [
        {
          "uuid": "89be9866-4807-6235-2a95-499cd23828df",
          "displayName": "Altitude Odin",
          "themeUuid": "537e0587-416c-f8f3-965c-bba3508156d7",
          "contentTierUuid": "0cebb8be-46d7-c12a-d306-e9907bfc5a25",
          "displayIcon": "https://media.valorant-api.com/weaponskins/89be9866-4807-6235-2a95-499cd23828df/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Airplane/HMG_Airplane_PrimaryAsset",
          "chromas": [
            {
              "uuid": "092a25a4-422f-f577-37ac-26a5d489c155",
              "displayName": "Altitude Odin",
              "displayIcon": null,
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/092a25a4-422f-f577-37ac-26a5d489c155/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Airplane/Chromas/Standard/HMG_Airplane_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "578e9077-4f88-260c-e54c-b988425c60e4",
              "displayName": "Altitude Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/578e9077-4f88-260c-e54c-b988425c60e4/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Airplane/Levels/HMG_Airplane_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "94c085e6-48e1-c879-2552-88bf7850c5a8",
          "displayName": "Xenohunter Odin",
          "themeUuid": "c251784d-40c3-0d8e-85f4-9da6ff3bc073",
          "contentTierUuid": "60bca009-4182-7998-dee7-b8a2558dc369",
          "displayIcon": "https://media.valorant-api.com/weaponskins/94c085e6-48e1-c879-2552-88bf7850c5a8/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Alien/HMG_Alien_PrimaryAsset",
          "chromas": [
            {
              "uuid": "3cc39f94-443e-e0e1-39be-21a9c9f3b0aa",
              "displayName": "Xenohunter Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/3cc39f94-443e-e0e1-39be-21a9c9f3b0aa/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/3cc39f94-443e-e0e1-39be-21a9c9f3b0aa/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Alien/Chromas/Standard/HMG_Alien_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "89baf0f4-4648-69fc-f0af-2fbc69b97b80",
              "displayName": "Xenohunter Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/89baf0f4-4648-69fc-f0af-2fbc69b97b80/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/32f7797f-4491-e21f-e40b-cfb639df3c97_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Alien/Levels/HMG_Alien_Lv1_PrimaryAsset"
            },
            {
              "uuid": "289d7473-472b-68da-9019-da9594a8a127",
              "displayName": "Xenohunter Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::HeartbeatAndMapSensor",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/bd9d37c4-4f6f-3766-c134-30b6782e714f_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Alien/Levels/HMG_Alien_Lv2_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "e02d6b2c-4e2e-a5f1-6bb8-38ac74485d7e",
          "displayName": "Tactiplay Odin",
          "themeUuid": "ba0640f1-41f0-849b-b16a-baaf7607b9b2",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/e02d6b2c-4e2e-a5f1-6bb8-38ac74485d7e/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Chibi/HMG_Chibi_PrimaryAsset",
          "chromas": [
            {
              "uuid": "1cefc3ab-4097-179d-79c3-a4a2bcdb4100",
              "displayName": "Tactiplay Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/1cefc3ab-4097-179d-79c3-a4a2bcdb4100/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/1cefc3ab-4097-179d-79c3-a4a2bcdb4100/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Chibi/Chromas/Standard/HMG_Chibi_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "ac60b469-450a-72fe-d2c0-67ae4da1baee",
              "displayName": "Tactiplay Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/ac60b469-450a-72fe-d2c0-67ae4da1baee/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Chibi/Levels/HMG_Chibi_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "abef8114-4495-6da7-2ade-02bd0f014fd3",
          "displayName": "Rune Stone Odin",
          "themeUuid": "65d8e156-4e4c-f9de-844f-3fa05a9624e4",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/abef8114-4495-6da7-2ade-02bd0f014fd3/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Codex/HMG_Codex_PrimaryAsset",
          "chromas": [
            {
              "uuid": "bf2fea95-4b73-05e3-9956-b8828a0f2edf",
              "displayName": "Rune Stone Odin",
              "displayIcon": null,
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/bf2fea95-4b73-05e3-9956-b8828a0f2edf/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Codex/Chromas/Standard/HMG_Codex_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "a00e2e2b-431a-1f84-e6d0-a58f4fe1e56b",
              "displayName": "Rune Stone Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/a00e2e2b-431a-1f84-e6d0-a58f4fe1e56b/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Codex/Levels/HMG_Codex_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "ba2592f7-45c8-8d25-1442-83971f3d95dd",
          "displayName": "Comet Odin",
          "themeUuid": "91640e14-4efe-e150-7b14-31a021012b3a",
          "contentTierUuid": "0cebb8be-46d7-c12a-d306-e9907bfc5a25",
          "displayIcon": "https://media.valorant-api.com/weaponskins/ba2592f7-45c8-8d25-1442-83971f3d95dd/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Comet/HMG_Comet_PrimaryAsset",
          "chromas": [
            {
              "uuid": "faba53e8-49c3-b167-1787-5d8caedbc552",
              "displayName": "Comet Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/faba53e8-49c3-b167-1787-5d8caedbc552/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/faba53e8-49c3-b167-1787-5d8caedbc552/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/faba53e8-49c3-b167-1787-5d8caedbc552/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Comet/Chromas/Standard/HMG_Comet_Standard_PrimaryAsset"
            },
            {
              "uuid": "563708e2-4895-f3ee-d44a-18911550d3b4",
              "displayName": "Comet Odin\r\n(Variant 1 Red)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/563708e2-4895-f3ee-d44a-18911550d3b4/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/563708e2-4895-f3ee-d44a-18911550d3b4/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/563708e2-4895-f3ee-d44a-18911550d3b4/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Comet/Chromas/v1/HMG_Comet_v1_PrimaryAsset"
            },
            {
              "uuid": "1c07da3a-47d9-97b7-f68f-c0b557d0ab13",
              "displayName": "Comet Odin\r\n(Variant 2 Pink)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/1c07da3a-47d9-97b7-f68f-c0b557d0ab13/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/1c07da3a-47d9-97b7-f68f-c0b557d0ab13/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/1c07da3a-47d9-97b7-f68f-c0b557d0ab13/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Comet/Chromas/v2/HMG_Comet_v2_PrimaryAsset"
            },
            {
              "uuid": "5b26ad8c-4382-f4a5-61c1-c0bce30d2bda",
              "displayName": "Comet Odin\r\n(Variant 3 Yellow)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/5b26ad8c-4382-f4a5-61c1-c0bce30d2bda/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/5b26ad8c-4382-f4a5-61c1-c0bce30d2bda/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/5b26ad8c-4382-f4a5-61c1-c0bce30d2bda/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Comet/Chromas/v3/HMG_Comet_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "5251cc6a-4ee9-9183-3b06-dcade2fc8cc7",
              "displayName": "Comet Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/5251cc6a-4ee9-9183-3b06-dcade2fc8cc7/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Comet/Levels/HMG_Comet_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "97af88e4-4176-9fa3-4a26-57919443dab7",
          "displayName": "Glitchpop Odin",
          "themeUuid": "5b014f36-414b-4703-9c65-1876c630feaa",
          "contentTierUuid": "e046854e-406c-37f4-6607-19a9ba8426fc",
          "displayIcon": "https://media.valorant-api.com/weaponskins/97af88e4-4176-9fa3-4a26-57919443dab7/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/HeavyMachineGun_Cyberpunk_PrimaryAsset",
          "chromas": [
            {
              "uuid": "9667983e-4c8c-e5b2-68d7-be84f9b3d46c",
              "displayName": "Glitchpop Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/9667983e-4c8c-e5b2-68d7-be84f9b3d46c/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/9667983e-4c8c-e5b2-68d7-be84f9b3d46c/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/9667983e-4c8c-e5b2-68d7-be84f9b3d46c/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/Chromas/Standard/HeavyMachineGun_Cyberpunk_Standard_PrimaryAsset"
            },
            {
              "uuid": "0b30b3e8-4696-7b7c-fed2-50b34234965a",
              "displayName": "Glitchpop Odin Level 4\r\n(Variant 1 Blue)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/0b30b3e8-4696-7b7c-fed2-50b34234965a/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/0b30b3e8-4696-7b7c-fed2-50b34234965a/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/0b30b3e8-4696-7b7c-fed2-50b34234965a/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/Chromas/Blue/HeavyMachineGun_Cyberpunk_Blue_PrimaryAsset"
            },
            {
              "uuid": "54caeb7f-4fc4-6adb-45a6-cfb6202d9c24",
              "displayName": "Glitchpop Odin Level 4\r\n(Variant 2 Red)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/54caeb7f-4fc4-6adb-45a6-cfb6202d9c24/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/54caeb7f-4fc4-6adb-45a6-cfb6202d9c24/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/54caeb7f-4fc4-6adb-45a6-cfb6202d9c24/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/Chromas/Red/HeavyMachineGun_Cyberpunk_Red_PrimaryAsset"
            },
            {
              "uuid": "bba7f46f-41ee-9e6c-c37a-dca8ee4bf50e",
              "displayName": "Glitchpop Odin Level 4\r\n(Variant 3 Gold)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/bba7f46f-41ee-9e6c-c37a-dca8ee4bf50e/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/bba7f46f-41ee-9e6c-c37a-dca8ee4bf50e/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/bba7f46f-41ee-9e6c-c37a-dca8ee4bf50e/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/Chromas/Gold/HeavyMachineGun_Cyberpunk_Gold_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "549b06bb-4704-25ce-19d5-c9b70b10de19",
              "displayName": "Glitchpop Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/549b06bb-4704-25ce-19d5-c9b70b10de19/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/72c8af91-f9f9-4044-801c-3e73ee2f2aa1_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/HeavyMachineGun_Cyberpunk_Lv1_PrimaryAsset"
            },
            {
              "uuid": "3e7d08f9-4067-1abe-8159-87b8e31fc463",
              "displayName": "Glitchpop Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::VFX",
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/3e7d08f9-4067-1abe-8159-87b8e31fc463/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/92c68b11-5e4c-48a7-acdd-32cd8dca7f01_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/HeavyMachineGun_Cyberpunk_Lv2_PrimaryAsset"
            },
            {
              "uuid": "ca91d540-4d2d-4946-70bd-97aae8117306",
              "displayName": "Glitchpop Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::VFX",
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/ca91d540-4d2d-4946-70bd-97aae8117306/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/18da6f4b-bb95-4158-9e8e-b785e4ba25d0_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/HeavyMachineGun_Cyberpunk_Lv3_PrimaryAsset"
            },
            {
              "uuid": "8c95559d-44fb-544d-00d7-8192ed38b17a",
              "displayName": "Glitchpop Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/8c95559d-44fb-544d-00d7-8192ed38b17a/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/1f8a6fe7-06a4-4cf4-8d1a-e4db58a0f700_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Cyberpunk/HeavyMachineGun_Cyberpunk_Lv4_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "2715f184-46cc-bec1-dd7c-e7b4d1aeb625",
          "displayName": "Nitro Odin",
          "themeUuid": "cc3f546b-410a-ddb6-6577-4aab0296496d",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": null,
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Electroflux2/HMG_Electroflux2_PrimaryAsset",
          "chromas": [
            {
              "uuid": "064eb4b1-477a-23ba-2f70-6a9732f68ece",
              "displayName": "Standard",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/064eb4b1-477a-23ba-2f70-6a9732f68ece/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/064eb4b1-477a-23ba-2f70-6a9732f68ece/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Electroflux2/Chromas/Standard/HMG_Electroflux2_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "5f86ac54-46b0-806f-f200-dd89fbcd7cb9",
              "displayName": "Nitro Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/5f86ac54-46b0-806f-f200-dd89fbcd7cb9/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Electroflux2/Levels/HMG_Electroflux2_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "85c76090-4de5-3a3a-a763-f4a7b779e8ed",
          "displayName": "Topotek Odin",
          "themeUuid": "418542b0-4f5c-3f37-4f5a-ecb1d3bc04ce",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/85c76090-4de5-3a3a-a763-f4a7b779e8ed/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Flyer/HMG_Flyer_PrimaryAsset",
          "chromas": [
            {
              "uuid": "2aa6e3af-424e-1278-e317-1faa895ab33c",
              "displayName": "Topotek Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/2aa6e3af-424e-1278-e317-1faa895ab33c/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/2aa6e3af-424e-1278-e317-1faa895ab33c/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Flyer/Chromas/Standard/HMG_Flyer_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "bdd376fa-42df-b1a3-dff5-df8c201ce53e",
              "displayName": "Topotek Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/bdd376fa-42df-b1a3-dff5-df8c201ce53e/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Flyer/Levels/HMG_Flyer_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "cda41b87-4d3a-c17c-5f6d-8990cc9c5efb",
          "displayName": ".EXE Odin",
          "themeUuid": "271874eb-491b-eae3-51f8-6f93f93035f9",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/cda41b87-4d3a-c17c-5f6d-8990cc9c5efb/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Grid/HeavyMachineGun_Grid_PrimaryAsset",
          "chromas": [
            {
              "uuid": "c1ed5bf3-4827-3e3a-ebbb-1ba42a226e59",
              "displayName": ".EXE Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/c1ed5bf3-4827-3e3a-ebbb-1ba42a226e59/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/c1ed5bf3-4827-3e3a-ebbb-1ba42a226e59/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Grid/Chromas/Standard/HeavyMachineGun_Grid_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "49292f21-4f5e-0a1a-3671-54b7ca8fe65a",
              "displayName": ".EXE Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/49292f21-4f5e-0a1a-3671-54b7ca8fe65a/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Grid/Levels/HeavyMachineGun_Grid_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "bd647d56-4542-19cd-e1ed-4fb429c78cf9",
          "displayName": "Neo Frontier Odin",
          "themeUuid": "3e8c23b5-47f3-fc77-788a-94852f7a111e",
          "contentTierUuid": "e046854e-406c-37f4-6607-19a9ba8426fc",
          "displayIcon": "https://media.valorant-api.com/weaponskins/bd647d56-4542-19cd-e1ed-4fb429c78cf9/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/HMG_Gunslinger_PrimaryAsset",
          "chromas": [
            {
              "uuid": "e20c028f-4e3e-0df7-551d-8c9eb51d08e7",
              "displayName": "Neo Frontier Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/e20c028f-4e3e-0df7-551d-8c9eb51d08e7/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/e20c028f-4e3e-0df7-551d-8c9eb51d08e7/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/e20c028f-4e3e-0df7-551d-8c9eb51d08e7/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Chromas/Standard/HMG_Gunslinger_Standard_PrimaryAsset"
            },
            {
              "uuid": "10be56f0-48eb-9293-3888-84ba8249cf46",
              "displayName": "Neo Frontier Odin Level 4\n(Variant 1 Gold)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/10be56f0-48eb-9293-3888-84ba8249cf46/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/10be56f0-48eb-9293-3888-84ba8249cf46/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/10be56f0-48eb-9293-3888-84ba8249cf46/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/9cc445f4-4446-1e1b-1a15-a3acdf74d52e_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Chromas/v1/HMG_Gunslinger_v1_PrimaryAsset"
            },
            {
              "uuid": "ce250c6a-4873-b305-4872-7392d0d84c1b",
              "displayName": "Neo Frontier Odin Level 4\n(Variant 2 Silver)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/ce250c6a-4873-b305-4872-7392d0d84c1b/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/ce250c6a-4873-b305-4872-7392d0d84c1b/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/ce250c6a-4873-b305-4872-7392d0d84c1b/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/9520c794-4614-636d-8b87-0d94cd5af43a_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Chromas/v2/HMG_Gunslinger_v2_PrimaryAsset"
            },
            {
              "uuid": "6de72b14-4680-8a4e-92bc-dd85a0aa67d7",
              "displayName": "Neo Frontier Odin Level 4\n(Variant 3 Copper)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/6de72b14-4680-8a4e-92bc-dd85a0aa67d7/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/6de72b14-4680-8a4e-92bc-dd85a0aa67d7/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/6de72b14-4680-8a4e-92bc-dd85a0aa67d7/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/5d99e67a-46b1-d041-e119-dfbf4b4ad1ac_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Chromas/v3/HMG_Gunslinger_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "6f60b3ce-4cbf-91ea-2ca7-59bbf1923751",
              "displayName": "Neo Frontier Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/6f60b3ce-4cbf-91ea-2ca7-59bbf1923751/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/729c9e7f-43be-dce4-b532-b99995902188_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Levels/HMG_Gunslinger_Lv1_PrimaryAsset"
            },
            {
              "uuid": "8c531daf-4b31-06f7-0cbe-f989029245d1",
              "displayName": "Neo Frontier Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::Animation",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/8a2d9288-4dd7-c902-d482-6c9684ceffc7_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Levels/HMG_Gunslinger_Lv2_PrimaryAsset"
            },
            {
              "uuid": "c5175815-41da-879b-c547-429ce4d7a191",
              "displayName": "Neo Frontier Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::Transformation",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/fe5acf44-4820-c780-b635-8587a6282432_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Levels/HMG_Gunslinger_Lv3_PrimaryAsset"
            },
            {
              "uuid": "220bfeea-44f4-51ee-c903-1dbfb1021ec4",
              "displayName": "Neo Frontier Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/bfe74105-40b6-b9e9-d614-00adfa9f53ce_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Gunslinger/Levels/HMG_Gunslinger_Lv4_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "157bcebe-484d-82e2-2a60-c8b4b11197ea",
          "displayName": "Prime//2.0 Odin",
          "themeUuid": "3264e5b6-4bd2-213b-eeab-4d8525dd4ffb",
          "contentTierUuid": "60bca009-4182-7998-dee7-b8a2558dc369",
          "displayIcon": "https://media.valorant-api.com/weaponskins/157bcebe-484d-82e2-2a60-c8b4b11197ea/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/HeavyMachineGun_HypeBeast2_PrimaryAsset",
          "chromas": [
            {
              "uuid": "9e59563c-4467-43df-3b58-2ca43c25abde",
              "displayName": "Prime//2.0 Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/9e59563c-4467-43df-3b58-2ca43c25abde/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/9e59563c-4467-43df-3b58-2ca43c25abde/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/9e59563c-4467-43df-3b58-2ca43c25abde/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Chromas/Standard/HeavyMachineGun_HypeBeast2_Standard_PrimaryAsset"
            },
            {
              "uuid": "ed34c641-4f1d-e38f-0018-4cb11fed9ee7",
              "displayName": "Prime//2.0 Odin Level 4\n(Variant 1 Gold)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/ed34c641-4f1d-e38f-0018-4cb11fed9ee7/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/ed34c641-4f1d-e38f-0018-4cb11fed9ee7/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/ed34c641-4f1d-e38f-0018-4cb11fed9ee7/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Chromas/Gold/HeavyMachineGun_HypeBeast2_Gold_PrimaryAsset"
            },
            {
              "uuid": "8617ebbb-418a-0819-1ca2-b383f8ae757c",
              "displayName": "Prime//2.0 Odin Level 4\n(Variant 2 Green)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/8617ebbb-418a-0819-1ca2-b383f8ae757c/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/8617ebbb-418a-0819-1ca2-b383f8ae757c/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/8617ebbb-418a-0819-1ca2-b383f8ae757c/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Chromas/Green/HeavyMachineGun_HypeBeast2_Green_PrimaryAsset"
            },
            {
              "uuid": "20ebb41a-4edc-0fa3-f3e9-b3b5d91e8524",
              "displayName": "Prime//2.0 Odin Level 4\n(Variant 3 Orange)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/20ebb41a-4edc-0fa3-f3e9-b3b5d91e8524/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/20ebb41a-4edc-0fa3-f3e9-b3b5d91e8524/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/20ebb41a-4edc-0fa3-f3e9-b3b5d91e8524/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Chromas/Orange/HeavyMachineGun_HypeBeast2_Orange_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "ef564ec3-497c-3038-543e-eb94bbe46437",
              "displayName": "Prime//2.0 Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/ef564ec3-497c-3038-543e-eb94bbe46437/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/07c47f26-4dd6-fe8f-bcdc-18b9c1425667_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Levels/HeavyMachineGun_HypeBeast2_Lv1_PrimaryAsset"
            },
            {
              "uuid": "ab04093d-489f-27f3-9e1b-1589db2185c8",
              "displayName": "Prime//2.0 Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::VFX",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/3ed8e0ec-48d9-5e12-58a7-9d9ea5e4fe4f_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Levels/HeavyMachineGun_HypeBeast2_Lv2_PrimaryAsset"
            },
            {
              "uuid": "7ca4b0b7-4ba5-9fc9-218c-0b8c865edaad",
              "displayName": "Prime//2.0 Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::Animation",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/f5aad4e7-437b-cb06-8cef-01b3dc04db72_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Levels/HeavyMachineGun_HypeBeast2_Lv3_PrimaryAsset"
            },
            {
              "uuid": "4f3eea45-4e3e-026f-66c9-658906a52d0b",
              "displayName": "Prime//2.0 Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/2ff2f95b-4199-a00a-08b7-df8a3a564d00_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/HypeBeast2/Levels/HeavyMachineGun_HypeBeast2_Lv4_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "72e724e9-4ba4-2d12-ce1a-8db1a528b9d3",
          "displayName": "Prism III Odin",
          "themeUuid": "0c2487bb-4cf9-78be-1bf1-c696f86b4aab",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/72e724e9-4ba4-2d12-ce1a-8db1a528b9d3/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Iridescence2/HMG_Iridescence2_PrimaryAsset",
          "chromas": [
            {
              "uuid": "b8bf7a9c-4724-a090-31b2-79b2211a6079",
              "displayName": "Prism III Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/b8bf7a9c-4724-a090-31b2-79b2211a6079/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/b8bf7a9c-4724-a090-31b2-79b2211a6079/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/b8bf7a9c-4724-a090-31b2-79b2211a6079/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Iridescence2/Chromas/Standard/HMG_Iridescence2_Standard_PrimaryAsset"
            },
            {
              "uuid": "0cdd55d5-44ba-d1fe-58ba-a881b50fdc0b",
              "displayName": "Prism III Odin\n(Variant 1 Pink)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/0cdd55d5-44ba-d1fe-58ba-a881b50fdc0b/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/0cdd55d5-44ba-d1fe-58ba-a881b50fdc0b/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/0cdd55d5-44ba-d1fe-58ba-a881b50fdc0b/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Iridescence2/Chromas/v1/HMG_Iridescence2_v1_PrimaryAsset"
            },
            {
              "uuid": "11c0fc56-4b49-84f4-8e86-9ca17b4d75b4",
              "displayName": "Prism III Odin\n(Variant 2 Green)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/11c0fc56-4b49-84f4-8e86-9ca17b4d75b4/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/11c0fc56-4b49-84f4-8e86-9ca17b4d75b4/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/11c0fc56-4b49-84f4-8e86-9ca17b4d75b4/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Iridescence2/Chromas/v2/HMG_Iridescence2_v2_PrimaryAsset"
            },
            {
              "uuid": "1aa0f35d-4a18-2054-d408-a18b2671507f",
              "displayName": "Prism III Odin\n(Variant 3 Purple)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/1aa0f35d-4a18-2054-d408-a18b2671507f/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/1aa0f35d-4a18-2054-d408-a18b2671507f/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/1aa0f35d-4a18-2054-d408-a18b2671507f/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Iridescence2/Chromas/v3/HMG_Iridescence2_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "7e323e82-44b6-1711-2028-10af32c632d5",
              "displayName": "Prism III Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/7e323e82-44b6-1711-2028-10af32c632d5/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Iridescence2/Levels/HMG_Iridescence2_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "9e648b20-4ed5-1f34-87a9-979cbe9a958a",
          "displayName": "Smite Odin",
          "themeUuid": "465336b0-4261-9dca-f6e8-109ad2c40381",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/9e648b20-4ed5-1f34-87a9-979cbe9a958a/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Lightning/HMG_Lightning_PrimaryAsset",
          "chromas": [
            {
              "uuid": "5da04039-4875-92ec-759f-5b9928d12b30",
              "displayName": "Smite Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/5da04039-4875-92ec-759f-5b9928d12b30/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/5da04039-4875-92ec-759f-5b9928d12b30/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Lightning/Chromas/Standard/HMG_Lightning_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "c59442a5-4b74-792c-d996-71a5fb340625",
              "displayName": "Smite Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/c59442a5-4b74-792c-d996-71a5fb340625/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Lightning/Levels/HMG_Lightning_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "5c13101a-45e4-d568-aade-d6b0dadedcd1",
          "displayName": "Coalition: Cobra Odin",
          "themeUuid": "4c6f01fd-4396-7d11-87b7-ac871a285241",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/5c13101a-45e4-d568-aade-d6b0dadedcd1/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Military/HMG_Military_PrimaryAsset",
          "chromas": [
            {
              "uuid": "e8874661-47d0-37cd-5f98-33b18e05872a",
              "displayName": "Coalition: Cobra Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/e8874661-47d0-37cd-5f98-33b18e05872a/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/e8874661-47d0-37cd-5f98-33b18e05872a/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/e8874661-47d0-37cd-5f98-33b18e05872a/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Military/Chromas/Standard/HMG_Military_Standard_PrimaryAsset"
            },
            {
              "uuid": "4a405c13-4440-94de-e894-82970565b71e",
              "displayName": "Coalition: Cobra Odin\n(Variant 1 Blue)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/4a405c13-4440-94de-e894-82970565b71e/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/4a405c13-4440-94de-e894-82970565b71e/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/4a405c13-4440-94de-e894-82970565b71e/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Military/Chromas/v1/HMG_Military_v1_PrimaryAsset"
            },
            {
              "uuid": "b49c9d6f-46c8-3689-01c1-bd81beb4050d",
              "displayName": "Coalition: Cobra Odin\n(Variant 2 Red)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/b49c9d6f-46c8-3689-01c1-bd81beb4050d/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/b49c9d6f-46c8-3689-01c1-bd81beb4050d/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/b49c9d6f-46c8-3689-01c1-bd81beb4050d/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Military/Chromas/v2/HMG_Military_v2_PrimaryAsset"
            },
            {
              "uuid": "c1e53748-4a97-c0a8-6759-59925fae909b",
              "displayName": "Coalition: Cobra Odin\n(Variant 3 Orange)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/c1e53748-4a97-c0a8-6759-59925fae909b/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/c1e53748-4a97-c0a8-6759-59925fae909b/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/c1e53748-4a97-c0a8-6759-59925fae909b/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Military/Chromas/v3/HMG_Military_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "dd2acfca-4366-7ac0-c13e-1eb2f1948273",
              "displayName": "Coalition: Cobra Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/dd2acfca-4366-7ac0-c13e-1eb2f1948273/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Military/Levels/HMG_Military_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "57523cf0-4574-968b-9f17-168e3bdb6d0d",
          "displayName": "Lightwave Odin",
          "themeUuid": "f9fb29bd-4592-ee49-0e2a-93a925d7332e",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/57523cf0-4574-968b-9f17-168e3bdb6d0d/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RGB/HMG_RGB_PrimaryAsset",
          "chromas": [
            {
              "uuid": "156158a5-4eb2-79ef-49e9-16a680fe93a9",
              "displayName": "Lightwave Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/156158a5-4eb2-79ef-49e9-16a680fe93a9/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/156158a5-4eb2-79ef-49e9-16a680fe93a9/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/156158a5-4eb2-79ef-49e9-16a680fe93a9/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RGB/Chromas/Standard/HMG_RGB_Standard_PrimaryAsset"
            },
            {
              "uuid": "a5a25319-4159-072f-5555-48824316ef00",
              "displayName": "Lightwave Odin\n(Variant 1 Blue)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/a5a25319-4159-072f-5555-48824316ef00/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/a5a25319-4159-072f-5555-48824316ef00/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/a5a25319-4159-072f-5555-48824316ef00/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RGB/Chromas/v1/HMG_RGB_v1_PrimaryAsset"
            },
            {
              "uuid": "4609dfdc-4d4e-463f-8fc9-c8a4c83369ce",
              "displayName": "Lightwave Odin\n(Variant 2 Red)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/4609dfdc-4d4e-463f-8fc9-c8a4c83369ce/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/4609dfdc-4d4e-463f-8fc9-c8a4c83369ce/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/4609dfdc-4d4e-463f-8fc9-c8a4c83369ce/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RGB/Chromas/v2/HMG_RGB_v2_PrimaryAsset"
            },
            {
              "uuid": "488f90c1-47e2-84dc-586d-609e53b5fe5b",
              "displayName": "Lightwave Odin\n(Variant 3 Gray)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/488f90c1-47e2-84dc-586d-609e53b5fe5b/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/488f90c1-47e2-84dc-586d-609e53b5fe5b/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/488f90c1-47e2-84dc-586d-609e53b5fe5b/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RGB/Chromas/v3/HMG_RGB_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "162beb92-4ab7-4383-da51-4b94ba90bd5d",
              "displayName": "Lightwave Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/162beb92-4ab7-4383-da51-4b94ba90bd5d/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RGB/Levels/HMG_RGB_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "65baa0cd-42ec-f99d-54a0-338d795b5824",
          "displayName": "Sensation Odin",
          "themeUuid": "548cc24d-4127-a9c3-060c-b2bada325474",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/65baa0cd-42ec-f99d-54a0-338d795b5824/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Rainbow/HMG_Rainbow_PrimaryAsset",
          "chromas": [
            {
              "uuid": "53ce97ed-47e0-4856-0446-f7bc86e869e0",
              "displayName": "Sensation Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/53ce97ed-47e0-4856-0446-f7bc86e869e0/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/53ce97ed-47e0-4856-0446-f7bc86e869e0/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Rainbow/Chromas/Standard/HMG_Rainbow_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "1bfea917-4262-37ba-3a76-04b937f2d0be",
              "displayName": "Sensation Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/1bfea917-4262-37ba-3a76-04b937f2d0be/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Rainbow/Levels/HMG_Rainbow_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "9c134f41-4c29-1bd8-682e-178e7f349c9b",
          "displayName": "Random Favorite Skin",
          "themeUuid": "0d7a5bfb-4850-098e-1821-d989bbfd58a8",
          "contentTierUuid": null,
          "displayIcon": "https://media.valorant-api.com/weaponskins/9c134f41-4c29-1bd8-682e-178e7f349c9b/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Random/HMG_Random_PrimaryAsset",
          "chromas": [
            {
              "uuid": "647f59fb-41ee-ed30-11f1-52b0814af441",
              "displayName": "Standard",
              "displayIcon": null,
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/647f59fb-41ee-ed30-11f1-52b0814af441/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Random/Chromas/Standard/HMG_Random_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "a89584fb-451d-fcce-0176-3bbaa6764a2e",
              "displayName": "Random Favorite Skin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/a89584fb-451d-fcce-0176-3bbaa6764a2e/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Random/Levels/HMG_Random_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "3bb7e1cd-4774-3b84-ab13-3fa8ca182f20",
          "displayName": "Orion Odin",
          "themeUuid": "5b8823bc-4419-c40c-7b62-19a3b7affd21",
          "contentTierUuid": "0cebb8be-46d7-c12a-d306-e9907bfc5a25",
          "displayIcon": "https://media.valorant-api.com/weaponskins/3bb7e1cd-4774-3b84-ab13-3fa8ca182f20/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RoboMecha/HMG_RoboMecha_PrimaryAsset",
          "chromas": [
            {
              "uuid": "2e9873e7-4fe8-e38f-76c8-2f8e30f626d6",
              "displayName": "Orion Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/2e9873e7-4fe8-e38f-76c8-2f8e30f626d6/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/2e9873e7-4fe8-e38f-76c8-2f8e30f626d6/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RoboMecha/Chromas/Standard/HMG_RoboMecha_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "10b34c7e-42cb-45bb-ec20-63b8b8332e6e",
              "displayName": "Orion Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/10b34c7e-42cb-45bb-ec20-63b8b8332e6e/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/RoboMecha/Levels/HMG_RoboMecha_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "67fb338a-4b21-ed70-7c2a-46bef4742b4f",
          "displayName": "Sentinels of Light Odin",
          "themeUuid": "7c85c5c1-4a47-eb3f-c81e-4693dd8b2a62",
          "contentTierUuid": "e046854e-406c-37f4-6607-19a9ba8426fc",
          "displayIcon": "https://media.valorant-api.com/weaponskins/67fb338a-4b21-ed70-7c2a-46bef4742b4f/displayicon.png",
          "wallpaper": "https://media.valorant-api.com/weaponskins/67fb338a-4b21-ed70-7c2a-46bef4742b4f/wallpaper.png",
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/HMG_SOL2_PrimaryAsset",
          "chromas": [
            {
              "uuid": "0b07deb9-4add-5c51-b381-3c8a1dd02cf2",
              "displayName": "Sentinels of Light Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/0b07deb9-4add-5c51-b381-3c8a1dd02cf2/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/0b07deb9-4add-5c51-b381-3c8a1dd02cf2/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/0b07deb9-4add-5c51-b381-3c8a1dd02cf2/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Chromas/Standard/HMG_SOL2_Standard_PrimaryAsset"
            },
            {
              "uuid": "a77bf207-4e2d-1a73-de52-2ebf528dbf6a",
              "displayName": "Sentinels of Light Odin Level 4\n(Variant 1 Pink)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/a77bf207-4e2d-1a73-de52-2ebf528dbf6a/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/a77bf207-4e2d-1a73-de52-2ebf528dbf6a/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/a77bf207-4e2d-1a73-de52-2ebf528dbf6a/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/e19a75a5-47ad-f10d-b860-29bf18e84541_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Chromas/v1/HMG_SOL2_v1_PrimaryAsset"
            },
            {
              "uuid": "d1738425-47ab-28f3-26c9-a3bb7a90c2c1",
              "displayName": "Sentinels of Light Odin Level 4\r\n(Variant 2 Red/Green)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/d1738425-47ab-28f3-26c9-a3bb7a90c2c1/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/d1738425-47ab-28f3-26c9-a3bb7a90c2c1/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/d1738425-47ab-28f3-26c9-a3bb7a90c2c1/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/9edd7ada-4875-633b-249f-d68b9a2a5320_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Chromas/v2/HMG_SOL2_v2_PrimaryAsset"
            },
            {
              "uuid": "c252d117-477c-138c-80e4-a2bab8d80c0c",
              "displayName": "Sentinels of Light Odin Level 4\r\n(Variant 3 Blue/Purple)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/c252d117-477c-138c-80e4-a2bab8d80c0c/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/c252d117-477c-138c-80e4-a2bab8d80c0c/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/c252d117-477c-138c-80e4-a2bab8d80c0c/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/452fc446-405a-96fb-6acd-99ab17b7a7ad_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Chromas/v3/HMG_SOL2_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "e96f3e84-471a-b20f-a3c8-39b3a607c9c8",
              "displayName": "Sentinels of Light Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/e96f3e84-471a-b20f-a3c8-39b3a607c9c8/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/de3742bb-488c-ef5b-8b06-e19ec538ec03_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Levels/HMG_SOL2_Lv1_PrimaryAsset"
            },
            {
              "uuid": "4e847ff2-46c6-d5fc-4b0c-b8ac66f08734",
              "displayName": "Sentinels of Light Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::SoundEffects",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/6bc05ff4-4c37-e7b8-3195-62bdb1115fbd_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Levels/HMG_SOL2_Lv2_PrimaryAsset"
            },
            {
              "uuid": "4961c412-433c-0996-c6b3-718957df2cc8",
              "displayName": "Sentinels of Light Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::Animation",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/368f7026-4488-60fe-8570-b8a91b0198ac_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Levels/HMG_SOL2_Lv3_PrimaryAsset"
            },
            {
              "uuid": "e8310f27-41e9-360a-55e7-1f887802ca68",
              "displayName": "Sentinels of Light Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/ef36d498-4aff-4c52-4083-9d88d89dc66d_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SOL2/Levels/HMG_SOL2_Lv4_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "02cce94a-4dc2-d11a-33cf-d8aba4e36202",
          "displayName": "Schema Odin",
          "themeUuid": "820b86ef-494c-d8f5-604b-07ae66c53c4c",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/02cce94a-4dc2-d11a-33cf-d8aba4e36202/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Schema/HMG_Schema_PrimaryAsset",
          "chromas": [
            {
              "uuid": "46afaf58-410c-661b-85d1-9eaafb4185b2",
              "displayName": "Schema Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/46afaf58-410c-661b-85d1-9eaafb4185b2/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/46afaf58-410c-661b-85d1-9eaafb4185b2/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Schema/Chromas/Standard/HMG_Schema_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "15a5516d-412f-2db0-6bb3-3cbe40a2355f",
              "displayName": "Schema Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/15a5516d-412f-2db0-6bb3-3cbe40a2355f/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Schema/Levels/HMG_Schema_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "468fdc95-443f-f1c2-bd22-fc8e1af0de39",
          "displayName": "Lycan\u0027s Bane Odin",
          "themeUuid": "ac3b20fe-423e-2f58-3baf-2cadeef47230",
          "contentTierUuid": "0cebb8be-46d7-c12a-d306-e9907bfc5a25",
          "displayIcon": "https://media.valorant-api.com/weaponskins/468fdc95-443f-f1c2-bd22-fc8e1af0de39/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SilverWolf/HMG_SilverWolf_PrimaryAsset",
          "chromas": [
            {
              "uuid": "42bdb11e-4d51-27eb-37e5-5c91c5c94a61",
              "displayName": "Lycan\u0027s Bane Odin",
              "displayIcon": null,
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/42bdb11e-4d51-27eb-37e5-5c91c5c94a61/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SilverWolf/Chromas/Standard/HMG_SilverWolf_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "e1a05cb4-4ac5-b15a-e86e-bab0f7e093ad",
              "displayName": "Lycan\u0027s Bane Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/e1a05cb4-4ac5-b15a-e86e-bab0f7e093ad/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/SilverWolf/Levels/HMG_SilverWolf_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "8dda01a6-4237-f430-ac70-c3ba677963e9",
          "displayName": "Reaver Odin",
          "themeUuid": "42110a3e-4363-50cc-5a1f-e4bc80f97916",
          "contentTierUuid": "60bca009-4182-7998-dee7-b8a2558dc369",
          "displayIcon": "https://media.valorant-api.com/weaponskins/8dda01a6-4237-f430-ac70-c3ba677963e9/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/HMG_Soulstealer2_PrimaryAsset",
          "chromas": [
            {
              "uuid": "cf42ad75-43db-5426-0645-a7a3fac452c5",
              "displayName": "Reaver Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/cf42ad75-43db-5426-0645-a7a3fac452c5/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/cf42ad75-43db-5426-0645-a7a3fac452c5/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/cf42ad75-43db-5426-0645-a7a3fac452c5/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Chromas/Standard/HMG_Soulstealer2_Standard_PrimaryAsset"
            },
            {
              "uuid": "0c643904-4cbc-8806-3198-eb86c22472bf",
              "displayName": "Reaver Odin Level 4\r\n(Variant 1 Red)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/0c643904-4cbc-8806-3198-eb86c22472bf/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/0c643904-4cbc-8806-3198-eb86c22472bf/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/0c643904-4cbc-8806-3198-eb86c22472bf/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/3a4dd823-4d9f-2532-3932-50a3ec401de9_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Chromas/v1/HMG_Soulstealer2_v1_PrimaryAsset"
            },
            {
              "uuid": "c8a7b6ae-4aaa-ba5b-9fad-0da1c8b25f9e",
              "displayName": "Reaver Odin Level 4\r\n(Variant 2 Black)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/c8a7b6ae-4aaa-ba5b-9fad-0da1c8b25f9e/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/c8a7b6ae-4aaa-ba5b-9fad-0da1c8b25f9e/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/c8a7b6ae-4aaa-ba5b-9fad-0da1c8b25f9e/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/c15396a1-4cec-a6f6-fc2f-eaa4dec7d389_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Chromas/v2/HMG_Soulstealer2_v2_PrimaryAsset"
            },
            {
              "uuid": "2f6c84b9-40dd-4c58-8f29-a6af6a31a09f",
              "displayName": "Reaver Odin Level 4\r\n(Variant 3 White)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/2f6c84b9-40dd-4c58-8f29-a6af6a31a09f/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/2f6c84b9-40dd-4c58-8f29-a6af6a31a09f/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/2f6c84b9-40dd-4c58-8f29-a6af6a31a09f/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/71a30e7e-46bd-d89a-d48d-3b9f2918912a_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Chromas/v3/HMG_Soulstealer2_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "f5ce6297-4cd4-4b09-3931-5f8b20a4317d",
              "displayName": "Reaver Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/f5ce6297-4cd4-4b09-3931-5f8b20a4317d/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/2ca99b17-4868-78b5-e68d-7aa0fd27a8fd_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Levels/HMG_Soulstealer2_Lv1_PrimaryAsset"
            },
            {
              "uuid": "2bb3b930-48d4-c32a-2972-6ba60f1e3d96",
              "displayName": "Reaver Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::VFX",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/df2bbfb2-4217-d364-98f2-deaca3b1edb7_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Levels/HMG_Soulstealer2_Lv2_PrimaryAsset"
            },
            {
              "uuid": "2a4c3d8a-44ec-0097-5afc-1bb1044605b4",
              "displayName": "Reaver Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::Animation",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/c2989dc2-4841-07eb-aa99-ce87bd7aafa2_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Levels/HMG_Soulstealer2_Lv3_PrimaryAsset"
            },
            {
              "uuid": "606a2388-440f-7858-e5e3-5a94d7aeb64f",
              "displayName": "Reaver Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/3edd48db-4518-e422-8cb6-869d243690a4_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Soulstealer2/Levels/HMG_Soulstealer2_Lv4_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "a7995818-409f-c79b-20b7-28ad642f3135",
          "displayName": "Sovereign Odin",
          "themeUuid": "3f183186-4895-f871-c4ab-75ad942baf8c",
          "contentTierUuid": "60bca009-4182-7998-dee7-b8a2558dc369",
          "displayIcon": "https://media.valorant-api.com/weaponskins/a7995818-409f-c79b-20b7-28ad642f3135/displayicon.png",
          "wallpaper": "https://media.valorant-api.com/weaponskins/a7995818-409f-c79b-20b7-28ad642f3135/wallpaper.png",
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/HMG_Sovereign2_PrimaryAsset",
          "chromas": [
            {
              "uuid": "a16360a6-4588-1cdf-d0e2-0ea215dc9faa",
              "displayName": "Sovereign Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/a16360a6-4588-1cdf-d0e2-0ea215dc9faa/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/a16360a6-4588-1cdf-d0e2-0ea215dc9faa/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/a16360a6-4588-1cdf-d0e2-0ea215dc9faa/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Chromas/Standard/HMG_Sovereign2_Standard_PrimaryAsset"
            },
            {
              "uuid": "4bc5a7a3-4283-687e-6a90-c3bd6ade055b",
              "displayName": "Sovereign Odin Level 4\n(Variant 1 Green)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/4bc5a7a3-4283-687e-6a90-c3bd6ade055b/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/4bc5a7a3-4283-687e-6a90-c3bd6ade055b/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/4bc5a7a3-4283-687e-6a90-c3bd6ade055b/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/b7fb6fd6-4a08-7d36-8383-c5a40591cad3_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Chromas/v1/HMG_Sovereign2_v1_PrimaryAsset"
            },
            {
              "uuid": "3d42de8a-4879-6ed6-451f-8faf531dcf54",
              "displayName": "Sovereign Odin Level 4\n(Variant 2 Teal)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/3d42de8a-4879-6ed6-451f-8faf531dcf54/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/3d42de8a-4879-6ed6-451f-8faf531dcf54/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/3d42de8a-4879-6ed6-451f-8faf531dcf54/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/3f48acde-4395-6f8d-57e7-2481b47b9953_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Chromas/v2/HMG_Sovereign2_v2_PrimaryAsset"
            },
            {
              "uuid": "d07a79a4-4dcd-8574-ead5-ba8d08786113",
              "displayName": "Sovereign Odin Level 4\n(Variant 3 Red)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/d07a79a4-4dcd-8574-ead5-ba8d08786113/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/d07a79a4-4dcd-8574-ead5-ba8d08786113/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/d07a79a4-4dcd-8574-ead5-ba8d08786113/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/fe55de9d-49e9-888d-6c24-7abcc3b83754_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Chromas/v3/HMG_Sovereign2_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "501811a4-4d60-d1fd-2775-ea92530ccbfa",
              "displayName": "Sovereign Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/501811a4-4d60-d1fd-2775-ea92530ccbfa/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/5f846174-4bf1-0ed6-a3c9-bf819ee434f8_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Levels/HMG_Sovereign2_Lv1_PrimaryAsset"
            },
            {
              "uuid": "ccdfeeab-419e-9b55-d8de-329ac2178d59",
              "displayName": "Sovereign Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::SoundEffects",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/052e3c95-433d-7690-27a7-f6a0d39ce286_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Levels/HMG_Sovereign2_Lv2_PrimaryAsset"
            },
            {
              "uuid": "d8ad2151-47b0-a16e-6b51-e1abc7c6bd23",
              "displayName": "Sovereign Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::VFX",
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/d8ad2151-47b0-a16e-6b51-e1abc7c6bd23/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/c9d2d5c2-4d4b-dd89-af98-53a290cf33ee_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Levels/HMG_Sovereign2_Lv3_PrimaryAsset"
            },
            {
              "uuid": "851475e6-4209-307b-1860-79a1363b81a1",
              "displayName": "Sovereign Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/5de58b7a-4b9b-7057-e708-53856165e222_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Sovereign2/Levels/HMG_Sovereign2_Lv4_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "f454efd1-49cb-372f-7096-d394df615308",
          "displayName": "Standard Odin",
          "themeUuid": "5a629df4-4765-0214-bd40-fbb96542941f",
          "contentTierUuid": null,
          "displayIcon": "https://media.valorant-api.com/weaponskins/f454efd1-49cb-372f-7096-d394df615308/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Standard/HeavyMachineGun_Standard_PrimaryAsset",
          "chromas": [
            {
              "uuid": "2f93861d-4b2f-2175-af0c-3ba0c736e257",
              "displayName": "Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/2f93861d-4b2f-2175-af0c-3ba0c736e257/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/2f93861d-4b2f-2175-af0c-3ba0c736e257/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Standard/Chromas/Standard/HeavyMachineGun_Standard_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "d91fb318-4e40-b4c9-8c0b-bb9da28bac55",
              "displayName": "Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/d91fb318-4e40-b4c9-8c0b-bb9da28bac55/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Standard/HeavyMachineGun_Standard_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "fa1c05fd-49fc-ad93-17d8-f0aaf11874cd",
          "displayName": "Evori Dreamwings Odin",
          "themeUuid": "8d9feaf5-4e2c-8ddb-1606-98868ad9e0c7",
          "contentTierUuid": "411e4a55-4e59-7757-41f0-86a53f101bb5",
          "displayIcon": "https://media.valorant-api.com/weaponskins/fa1c05fd-49fc-ad93-17d8-f0aaf11874cd/displayicon.png",
          "wallpaper": "https://media.valorant-api.com/weaponskins/fa1c05fd-49fc-ad93-17d8-f0aaf11874cd/wallpaper.png",
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/HMG_StarPower_PrimaryAsset",
          "chromas": [
            {
              "uuid": "6e1ef3fb-4197-87f8-ecfd-98809b0575ec",
              "displayName": "Evori Dreamwings Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/6e1ef3fb-4197-87f8-ecfd-98809b0575ec/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/6e1ef3fb-4197-87f8-ecfd-98809b0575ec/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/6e1ef3fb-4197-87f8-ecfd-98809b0575ec/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Chromas/Standard/HMG_StarPower_Standard_PrimaryAsset"
            },
            {
              "uuid": "11f0b0ed-4152-3a57-78cb-71a5765adb6f",
              "displayName": "Evori Dreamwings Odin Level 4\r\n(Variant 1 Blue)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/11f0b0ed-4152-3a57-78cb-71a5765adb6f/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/11f0b0ed-4152-3a57-78cb-71a5765adb6f/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/11f0b0ed-4152-3a57-78cb-71a5765adb6f/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/6660e40a-4563-52e6-89a6-5582c3b64b3a_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Chromas/v1/HMG_StarPower_v1_PrimaryAsset"
            },
            {
              "uuid": "f6581e76-4be5-c878-04a9-688068b76dbb",
              "displayName": "Evori Dreamwings Odin Level 4\r\n(Variant 2 Green)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/f6581e76-4be5-c878-04a9-688068b76dbb/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/f6581e76-4be5-c878-04a9-688068b76dbb/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/f6581e76-4be5-c878-04a9-688068b76dbb/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/2799d005-4d82-a18a-884e-e898eafd9e0a_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Chromas/v2/HMG_StarPower_v2_PrimaryAsset"
            },
            {
              "uuid": "85bf962d-43f7-bcee-e2dc-1aa388186b4d",
              "displayName": "Evori Dreamwings Odin Level 4\r\n(Variant 3 Pink)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/85bf962d-43f7-bcee-e2dc-1aa388186b4d/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/85bf962d-43f7-bcee-e2dc-1aa388186b4d/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/85bf962d-43f7-bcee-e2dc-1aa388186b4d/swatch.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/138c0343-441b-fb67-5f73-36a06d70fee9_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Chromas/v3/HMG_StarPower_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "e6b64b6e-41cb-a99f-dceb-c0bc95b93d0d",
              "displayName": "Evori Dreamwings Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/e6b64b6e-41cb-a99f-dceb-c0bc95b93d0d/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/c4856f88-4097-2486-ef3b-93a5d2e0bf52_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Levels/HMG_StarPower_Lv1_PrimaryAsset"
            },
            {
              "uuid": "ec3d8ba4-4c69-128f-6344-86bfbfb88cae",
              "displayName": "Evori Dreamwings Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::SoundEffects",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/b8bda705-4478-3976-b6b4-5da91991c8c0_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Levels/HMG_StarPower_Lv2_PrimaryAsset"
            },
            {
              "uuid": "e80398c5-4dd1-daeb-824b-a79c57c8540d",
              "displayName": "Evori Dreamwings Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::Animation",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/31213d67-4e1b-917b-f486-93958ac7a688_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Levels/HMG_StarPower_Lv3_PrimaryAsset"
            },
            {
              "uuid": "454354b7-49ea-550c-84de-cc97781fe0d7",
              "displayName": "Evori Dreamwings Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": null,
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/faca3515-4365-003c-1355-679be0606088_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StarPower/Levels/HMG_StarPower_Lv4_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "14796249-4f23-9d52-4ea1-d8878099c01e",
          "displayName": "Freehand Odin",
          "themeUuid": "51b0a893-40fb-acb4-3214-be89dac1983a",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/14796249-4f23-9d52-4ea1-d8878099c01e/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StickerPop/HMG_StickerPop_PrimaryAsset",
          "chromas": [
            {
              "uuid": "818aa4aa-4e44-606d-bfe7-3abf3fca62ea",
              "displayName": "Freehand Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/818aa4aa-4e44-606d-bfe7-3abf3fca62ea/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/818aa4aa-4e44-606d-bfe7-3abf3fca62ea/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StickerPop/Chromas/Standard/HMG_StickerPop_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "7d9adf46-48c5-9247-7121-1ba9e328e70c",
              "displayName": "Freehand Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/7d9adf46-48c5-9247-7121-1ba9e328e70c/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StickerPop/Levels/HMG_StickerPop_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "befa2f32-410f-a418-d8d3-b194dcf2ec6d",
          "displayName": "Aerosol Odin",
          "themeUuid": "625f8165-44d8-0104-33da-e38da60cdb11",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/befa2f32-410f-a418-d8d3-b194dcf2ec6d/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StreetArt/HMG_StreetArt_PrimaryAsset",
          "chromas": [
            {
              "uuid": "309941f1-48b5-edc2-5e07-eaa7a435fedc",
              "displayName": "Aerosol Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/309941f1-48b5-edc2-5e07-eaa7a435fedc/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/309941f1-48b5-edc2-5e07-eaa7a435fedc/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StreetArt/Chromas/Standard/HMG_StreetArt_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "657b3476-4919-7b4b-be54-aeaedad999e6",
              "displayName": "Aerosol Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/657b3476-4919-7b4b-be54-aeaedad999e6/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/StreetArt/Levels/HMG_StreetArt_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "14f9d94a-4add-10a9-588d-48a7111da96f",
          "displayName": "Fortune\u0027s Hand Odin",
          "themeUuid": "70b4ccb1-4508-3a83-b75d-669f76c11e88",
          "contentTierUuid": "12683d76-48d7-84a3-4e09-6985794f0445",
          "displayIcon": "https://media.valorant-api.com/weaponskins/14f9d94a-4add-10a9-588d-48a7111da96f/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Tarot/HMG_Tarot_PrimaryAsset",
          "chromas": [
            {
              "uuid": "60118ac3-429b-7186-5e5d-07af412b2e9a",
              "displayName": "Fortune\u0027s Hand Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/60118ac3-429b-7186-5e5d-07af412b2e9a/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/60118ac3-429b-7186-5e5d-07af412b2e9a/fullrender.png",
              "swatch": null,
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Tarot/Chromas/Standard/HMG_Tarot_Standard_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "1fa403bd-4faa-b3c3-d7ce-f9bfc3ba45b7",
              "displayName": "Fortune\u0027s Hand Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/1fa403bd-4faa-b3c3-d7ce-f9bfc3ba45b7/displayicon.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Tarot/Levels/HMG_Tarot_Lv1_PrimaryAsset"
            }
          ]
        },
        {
          "uuid": "85ed3f9d-4e59-a709-8faf-bc86effb3a07",
          "displayName": "BlastX Odin",
          "themeUuid": "40a12450-468d-4632-8374-3abd1d4d5eb9",
          "contentTierUuid": "e046854e-406c-37f4-6607-19a9ba8426fc",
          "displayIcon": "https://media.valorant-api.com/weaponskins/85ed3f9d-4e59-a709-8faf-bc86effb3a07/displayicon.png",
          "wallpaper": null,
          "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/HeavyMachineGun_URF_PrimaryAsset",
          "chromas": [
            {
              "uuid": "f55aac92-4420-7038-634a-8fb3fc9a936d",
              "displayName": "BlastX Odin",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/f55aac92-4420-7038-634a-8fb3fc9a936d/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/f55aac92-4420-7038-634a-8fb3fc9a936d/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/f55aac92-4420-7038-634a-8fb3fc9a936d/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Chromas/Standard/HeavyMachineGun_URF_Standard_PrimaryAsset"
            },
            {
              "uuid": "3b809ee0-4af3-b604-9ac4-a799241289e7",
              "displayName": "BlastX Odin Level 4\n(Variant 1 Black)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/3b809ee0-4af3-b604-9ac4-a799241289e7/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/3b809ee0-4af3-b604-9ac4-a799241289e7/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/3b809ee0-4af3-b604-9ac4-a799241289e7/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Chromas/v1/HeavyMachineGun_URF_v1_PrimaryAsset"
            },
            {
              "uuid": "027dda2f-491d-e55e-1d32-acb7f55008f9",
              "displayName": "BlastX Odin Level 4\n(Variant 2 Yellow)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/027dda2f-491d-e55e-1d32-acb7f55008f9/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/027dda2f-491d-e55e-1d32-acb7f55008f9/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/027dda2f-491d-e55e-1d32-acb7f55008f9/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Chromas/v2/HeavyMachineGun_URF_v2_PrimaryAsset"
            },
            {
              "uuid": "e2c992bc-4099-eeac-9e3f-2e9d6a6bf757",
              "displayName": "BlastX Odin Level 4\n(Variant 3 Pink)",
              "displayIcon": "https://media.valorant-api.com/weaponskinchromas/e2c992bc-4099-eeac-9e3f-2e9d6a6bf757/displayicon.png",
              "fullRender": "https://media.valorant-api.com/weaponskinchromas/e2c992bc-4099-eeac-9e3f-2e9d6a6bf757/fullrender.png",
              "swatch": "https://media.valorant-api.com/weaponskinchromas/e2c992bc-4099-eeac-9e3f-2e9d6a6bf757/swatch.png",
              "streamedVideo": null,
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Chromas/v3/HeavyMachineGun_URF_v3_PrimaryAsset"
            }
          ],
          "levels": [
            {
              "uuid": "5123ed96-48de-aeab-eaea-2b8ea689f5a2",
              "displayName": "BlastX Odin",
              "levelItem": null,
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/5123ed96-48de-aeab-eaea-2b8ea689f5a2/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/943c425e-4b3d-0d29-15d7-8d88d95b886e_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Levels/HeavyMachineGun_URF_Lv1_PrimaryAsset"
            },
            {
              "uuid": "35e89d1a-477c-3d54-82aa-3599a832fe0b",
              "displayName": "BlastX Odin Level 2",
              "levelItem": "EEquippableSkinLevelItem::VFX",
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/35e89d1a-477c-3d54-82aa-3599a832fe0b/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/3d8cc4ea-4de2-4863-4c9c-c3b86cc6884c_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Levels/HeavyMachineGun_URF_Lv2_PrimaryAsset"
            },
            {
              "uuid": "31201e02-4873-a6db-1ce6-6e93a620370c",
              "displayName": "BlastX Odin Level 3",
              "levelItem": "EEquippableSkinLevelItem::Animation",
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/31201e02-4873-a6db-1ce6-6e93a620370c/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/c3008423-4a36-7d58-51f8-81a2eea188c3_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Levels/HeavyMachineGun_URF_Lv3_PrimaryAsset"
            },
            {
              "uuid": "7637fff1-4275-590a-6181-be95eae8786a",
              "displayName": "BlastX Odin Level 4",
              "levelItem": "EEquippableSkinLevelItem::Finisher",
              "displayIcon": "https://media.valorant-api.com/weaponskinlevels/7637fff1-4275-590a-6181-be95eae8786a/displayicon.png",
              "streamedVideo": "https://valorant.dyn.riotcdn.net/x/videos/release-09.04/463b382a-4a76-4eb3-a699-bc8e856335c4_default_universal.mp4",
              "assetPath": "ShooterGame/Content/Equippables/Guns/HvyMachineGuns/HMG/Urf/Levels/HeavyMachineGun_URF_Lv4_PrimaryAsset"
            }
          ]
        }
      ]
    }`)

	return json
}
