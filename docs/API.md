FORMAT: 1A
HOST: https://zsr.octane.gg

# Octane ZSR - Rocket League Esports API

This API powers Octane.gg and is available for public use. If you're actively using the API, be sure to join our [Discord](http://discord.gg/gxHfxWq) and let us know what you're working on!

# Group Events

## Events [/events]

### List Events [GET]

+ Parameters
    + name: RLCS (string, optional) - a portion of the event name
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter events before this date
    
    + after: `2016-12-03` (date, optional) - filter events after this date
    
    + date: `2016-12-03` (date, optional) - filter events on this date
    
    + sort: name (string, optional) - field to sort by
    
    + order: `asc` (enum[string], optional) - order of sort
        + Members
          + asc
          + desc
    
    + page: 1 (number) - page number
        + Default: 1
    
    + perPage: 20 (number) - results per page
        + Default: 50

* Response 200 (application/json)
  {"events":[{"\_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","startDate":"2016-12-03T00:00:00Z","endDate":"2016-12-04T00:00:00Z","region":"INT","mode":3,"prize":{"amount":125000,"currency":"USD"},"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"],"stages":[{"_id":0,"name":"Main Event","format":"bracket-8de","region":"INT","startDate":"2016-12-03T00:00:00Z","endDate":"2016-12-04T00:00:00Z","prize":{"amount":125000,"currency":"USD"},"liquipedia":"https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/Season_2","location":{"venue":"Theater Amsterdam","city":"Amsterdam","country":"nl"}}]}]}

## Event [/events/{id}]

- Parameters
  + id: 5f35882d53fbbb5894b43040 (required, string) - an event id

### Get Event [GET]

- Response 200 (application/json)
  {"\_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","startDate":"2016-12-03T00:00:00Z","endDate":"2016-12-04T00:00:00Z","region":"INT","mode":3,"prize":{"amount":125000,"currency":"USD"},"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"],"stages":[{"_id":0,"name":"Main Event","format":"bracket-8de","region":"INT","startDate":"2016-12-03T00:00:00Z","endDate":"2016-12-04T00:00:00Z","prize":{"amount":125000,"currency":"USD"},"liquipedia":"https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/Season_2","location":{"venue":"Theater Amsterdam","city":"Amsterdam","country":"nl"}}]}

## Event [/events/{id}/matches]

- Parameters
  + id: 5f35882d53fbbb5894b43040 (required, string) - an event id

    + stage: 1 (number, optional) - a stage id

### Get Event Matches [GET]

- Response 200 (application/json)
  {"matches":[{"_id":"6043152fa09e7fba40d2ae62","slug":"ae62-flipsid3-tactics-vs-nrg-esports","octane_id":"0350109","event":{"_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","region":"INT","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"]},"stage":{"_id":0,"name":"Main Event","format":"bracket-8de"},"date":"2016-12-04T00:00:00Z","format":{"type":"best","length":5},"blue":{"score":3,"winner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"},"stats":{"core":{"shots":34,"goals":6,"saves":14,"assists":5,"score":2620,"shootingPercentage":17.647058823529413},"boost":{"bpm":4841,"bcpm":4951.03124,"avgAmount":600.61,"amountCollected":26736,"amountStolen":5321,"amountCollectedBig":19445,"amountStolenBig":3178,"amountCollectedSmall":7291,"amountStolenSmall":2143,"countCollectedBig":225,"countStolenBig":39,"countCollectedSmall":653,"countStolenSmall":184,"amountOverfill":3162,"amountOverfillStolen":647,"amountUsedWhileSupersonic":3745,"timeZeroBoost":420.44,"timeFullBoost":418.98,"timeBoost0To25":1358.4,"timeBoost25To50":923.6800000000001,"timeBoost50To75":754.07,"timeBoost75To100":891.18},"movement":{"totalDistance":5828633,"timeSupersonicSpeed":551.12,"timeBoostSpeed":1636.22,"timeSlowSpeed":1803.4099999999999,"timeGround":2299.33,"timeLowAir":1528.4899999999998,"timeHighAir":162.89999999999998,"timePowerslide":90.80999999999999,"countPowerslide":770},"positioning":{"timeDefensiveThird":1820.74,"timeNeutralThird":1319.62,"timeOffensiveThird":850.3599999999999,"timeDefensiveHalf":2525.01,"timeOffensiveHalf":1465.1799999999998,"timeBehindBall":2979.85,"timeInfrontBall":1010.86},"demo":{"inflicted":6,"taken":8}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","country":"it"},"stats":{"core":{"shots":9,"goals":3,"saves":5,"assists":2,"score":940,"shootingPercentage":33.33333333333333},"boost":{"bpm":1355,"bcpm":1383.3523500000001,"avgAmount":209.42000000000002,"amountCollected":7488,"amountStolen":1478,"amountCollectedBig":4925,"amountStolenBig":814,"amountCollectedSmall":2563,"amountStolenSmall":664,"countCollectedBig":60,"countStolenBig":11,"countCollectedSmall":225,"countStolenSmall":56,"amountOverfill":1048,"amountOverfillStolen":196,"amountUsedWhileSupersonic":889,"timeZeroBoost":73.11,"percentZeroBoost":5.66867,"timeFullBoost":103.88999999999999,"percentFullBoost":7.945092724999999,"timeBoost0To25":371.58000000000004,"timeBoost25To50":339.95000000000005,"timeBoost50To75":302.69,"timeBoost75To100":293.02,"percentBoost0To25":28.43400725,"percentBoost25To50":26.14506925,"percentBoost50To75":23.053028750000003,"percentBoost75To100":22.3678945},"movement":{"avgSpeed":5746,"totalDistance":1836792,"timeSupersonicSpeed":132.97,"timeBoostSpeed":522.86,"timeSlowSpeed":677.12,"timeGround":769.29,"timeLowAir":515.0999999999999,"timeHighAir":48.55,"timePowerslide":29.229999999999997,"countPowerslide":228,"avgPowerslideDuration":0.51,"avgSpeedPercentage":62.456519,"percentSlowSpeed":50.90838250000001,"percentBoostSpeed":39.1156835,"percentSupersonicSpeed":9.97593335,"percentGround":57.59838475,"percentLowAir":38.7270215,"percentHighAir":3.67459345},"positioning":{"avgDistanceToBall":12416,"avgDistanceToBallPossession":12348,"avgDistanceToBallNoPossession":12606,"avgDistanceToMates":14960,"timeDefensiveThird":644.42,"timeNeutralThird":447.56,"timeOffensiveThird":240.95999999999998,"timeDefensiveHalf":893.57,"timeOffensiveHalf":439.23,"timeBehindBall":1049.42,"timeInfrontBall":283.51000000000005,"timeMostBack":518,"timeMostForward":331.6,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":418.9,"timeFarthestFromBall":425.80000000000007,"percentDefensiveThird":48.467209999999994,"percentOffensiveThird":18.10444525,"percentNeutralThird":33.42834875,"percentDefensiveHalf":67.11758375,"percentOffensiveHalf":32.882416,"percentBehindBall":78.83672,"percentInfrontBall":21.16328125,"percentMostBack":40.0437585,"percentMostForward":25.44998075,"percentClosestToBall":32.311219,"percentFarthestFromBall":32.8769945},"demo":{"inflicted":4,"taken":2}},"advanced":{"goalParticipation":83.33333333333334,"rating":1.0904344687742749}},{"player":{"_id":"6082fb1d0d9dcf9da5a4d079","slug":"d079-greazymeister","tag":"gReazymeister"},"stats":{"core":{"shots":11,"goals":1,"saves":6,"assists":1,"score":855,"shootingPercentage":9.090909090909092},"boost":{"bpm":1749,"bcpm":1780.08505,"avgAmount":191.92000000000002,"amountCollected":9610,"amountStolen":1666,"amountCollectedBig":7141,"amountStolenBig":951,"amountCollectedSmall":2469,"amountStolenSmall":715,"countCollectedBig":80,"countStolenBig":11,"countCollectedSmall":224,"countStolenSmall":59,"amountOverfill":925,"amountOverfillStolen":145,"amountUsedWhileSupersonic":1232,"timeZeroBoost":193.96,"percentZeroBoost":15.053702249999999,"timeFullBoost":118.42999999999999,"percentFullBoost":9.116349875000001,"timeBoost0To25":543.5799999999999,"timeBoost25To50":265.74,"timeBoost50To75":222.40000000000003,"timeBoost75To100":286.5,"percentBoost0To25":41.26221275,"percentBoost25To50":20.05843075,"percentBoost50To75":16.95397875,"percentBoost75To100":21.72537775},"movement":{"avgSpeed":6211,"totalDistance":1977383,"timeSupersonicSpeed":191.99,"timeBoostSpeed":568.89,"timeSlowSpeed":567.1500000000001,"timeGround":788.93,"timeLowAir":481.87,"timeHighAir":57.209999999999994,"timePowerslide":30.569999999999993,"countPowerslide":235,"avgPowerslideDuration":0.52,"avgSpeedPercentage":67.51087150000001,"percentSlowSpeed":42.834596499999996,"percentBoostSpeed":42.83116575,"percentSupersonicSpeed":14.33424075,"percentGround":59.40004425,"percentLowAir":36.317852,"percentHighAir":4.282102275},"positioning":{"avgDistanceToBall":12676,"avgDistanceToBallPossession":12163,"avgDistanceToBallNoPossession":13231,"avgDistanceToMates":15338,"timeDefensiveThird":633.36,"timeNeutralThird":431.29,"timeOffensiveThird":263.37,"timeDefensiveHalf":869.01,"timeOffensiveHalf":458.75,"timeBehindBall":1001.28,"timeInfrontBall":326.74,"timeMostBack":480.20000000000005,"timeMostForward":452.7,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":449.4,"timeFarthestFromBall":436.9,"percentDefensiveThird":47.78079675,"percentOffensiveThird":19.832879125,"percentNeutralThird":32.386322,"percentDefensiveHalf":65.47910325,"percentOffensiveHalf":34.520901,"percentBehindBall":75.44333999999999,"percentInfrontBall":24.556661000000002,"percentMostBack":37.11360125,"percentMostForward":34.8868515,"percentClosestToBall":34.765083749999995,"percentFarthestFromBall":33.646304},"demo":{"inflicted":2,"taken":3}},"advanced":{"goalParticipation":33.33333333333333,"rating":0.6576622683574388}},{"player":{"_id":"5f3d8fdd95f40596eae23d98","slug":"3d98-markydooda","tag":"Markydooda","country":"ab"},"stats":{"core":{"shots":14,"goals":2,"saves":3,"assists":2,"score":825,"shootingPercentage":14.285714285714285},"boost":{"bpm":1737,"bcpm":1787.59384,"avgAmount":199.27,"amountCollected":9638,"amountStolen":2177,"amountCollectedBig":7379,"amountStolenBig":1413,"amountCollectedSmall":2259,"amountStolenSmall":764,"countCollectedBig":85,"countStolenBig":17,"countCollectedSmall":204,"countStolenSmall":69,"amountOverfill":1189,"amountOverfillStolen":306,"amountUsedWhileSupersonic":1624,"timeZeroBoost":153.37,"percentZeroBoost":11.839296999999998,"timeFullBoost":196.66000000000003,"percentFullBoost":15.160770750000001,"timeBoost0To25":443.24,"timeBoost25To50":317.99,"timeBoost50To75":228.98000000000002,"timeBoost75To100":311.65999999999997,"percentBoost0To25":34.12572475,"percentBoost25To50":24.450708749999997,"percentBoost50To75":17.511339625,"percentBoost75To100":23.91222475},"movement":{"avgSpeed":6325,"totalDistance":2014458,"timeSupersonicSpeed":226.16,"timeBoostSpeed":544.47,"timeSlowSpeed":559.14,"timeGround":741.1099999999999,"timeLowAir":531.52,"timeHighAir":57.14,"timePowerslide":31.01,"countPowerslide":307,"avgPowerslideDuration":0.4,"avgSpeedPercentage":68.7500005,"percentSlowSpeed":41.99217525,"percentBoostSpeed":40.93057,"percentSupersonicSpeed":17.0772575,"percentGround":55.656828749999995,"percentLowAir":40.04468875,"percentHighAir":4.298480475},"positioning":{"avgDistanceToBall":12663,"avgDistanceToBallPossession":12355,"avgDistanceToBallNoPossession":12989,"avgDistanceToMates":15613,"timeDefensiveThird":542.96,"timeNeutralThird":440.77,"timeOffensiveThird":346.03,"timeDefensiveHalf":762.4300000000001,"timeOffensiveHalf":567.1999999999999,"timeBehindBall":929.1500000000001,"timeInfrontBall":400.61,"timeMostBack":331.1,"timeMostForward":523.9,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":440.30000000000007,"timeFarthestFromBall":466.5,"percentDefensiveThird":40.66040425,"percentOffensiveThird":26.109765250000002,"percentNeutralThird":33.2298305,"percentDefensiveHalf":57.213438999999994,"percentOffensiveHalf":42.78656325,"percentBehindBall":69.843096,"percentInfrontBall":30.156903500000002,"percentMostBack":25.40523275,"percentMostForward":40.625619,"percentClosestToBall":33.915815,"percentFarthestFromBall":36.031493499999996},"demo":{"inflicted":0,"taken":3}},"advanced":{"goalParticipation":66.66666666666666,"rating":0.9047666428357998}}]},"orange":{"score":1,"team":{"team":{"_id":"6020bc70f1e4807cc70023a0","slug":"23a0-nrg-esports","name":"NRG Esports","image":"https://griffon.octane.gg/teams/nrg-esports.png"},"stats":{"core":{"shots":22,"goals":5,"saves":22,"assists":3,"score":2540,"shootingPercentage":22.727272727272727},"boost":{"bpm":5095,"bcpm":5181.857169999999,"avgAmount":568.9,"amountCollected":27984,"amountStolen":6409,"amountCollectedBig":20642,"amountStolenBig":4724,"amountCollectedSmall":7342,"amountStolenSmall":1685,"countCollectedBig":237,"countStolenBig":53,"countCollectedSmall":632,"countStolenSmall":148,"amountOverfill":2954,"amountOverfillStolen":547,"amountUsedWhileSupersonic":5117,"timeZeroBoost":639.2099999999999,"timeFullBoost":445.38,"timeBoost0To25":1479.33,"timeBoost25To50":829.08,"timeBoost50To75":667.6700000000001,"timeBoost75To100":919.11},"movement":{"totalDistance":5923761,"timeSupersonicSpeed":595.7900000000001,"timeBoostSpeed":1719.2600000000002,"timeSlowSpeed":1681.8899999999999,"timeGround":2351.83,"timeLowAir":1494.1999999999998,"timeHighAir":150.91,"timePowerslide":73.97999999999999,"countPowerslide":550},"positioning":{"timeDefensiveThird":2032.74,"timeNeutralThird":1183.71,"timeOffensiveThird":780.46,"timeDefensiveHalf":2676.79,"timeOffensiveHalf":1319.5300000000002,"timeBehindBall":2741.89,"timeInfrontBall":1255.06},"demo":{"inflicted":8,"taken":6}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7a","slug":"3d7a-fireburner","tag":"Fireburner","country":"us"},"stats":{"core":{"shots":8,"goals":1,"saves":9,"assists":0,"score":775,"shootingPercentage":12.5},"boost":{"bpm":1744,"bcpm":1748.36873,"avgAmount":195.23,"amountCollected":9449,"amountStolen":2173,"amountCollectedBig":7416,"amountStolenBig":1652,"amountCollectedSmall":2033,"amountStolenSmall":521,"countCollectedBig":85,"countStolenBig":19,"countCollectedSmall":186,"countStolenSmall":51,"amountOverfill":1070,"amountOverfillStolen":220,"amountUsedWhileSupersonic":1742,"timeZeroBoost":156.04,"percentZeroBoost":12.079158625000002,"timeFullBoost":169.47,"percentFullBoost":12.9679785,"timeBoost0To25":455.68999999999994,"timeBoost25To50":287.96999999999997,"timeBoost50To75":269.03999999999996,"timeBoost75To100":290.03,"percentBoost0To25":34.95103375,"percentBoost25To50":22.17023925,"percentBoost50To75":20.63573975,"percentBoost75To100":22.24298725},"movement":{"avgSpeed":6069,"totalDistance":1938176,"timeSupersonicSpeed":212.43,"timeBoostSpeed":528.79,"timeSlowSpeed":591.11,"timeGround":792.7,"timeLowAir":513.06,"timeHighAir":26.560000000000002,"timePowerslide":21.509999999999998,"countPowerslide":164,"avgPowerslideDuration":0.52,"avgSpeedPercentage":65.967393,"percentSlowSpeed":44.25205375,"percentBoostSpeed":39.8215655,"percentSupersonicSpeed":15.926380625,"percentGround":59.421816500000006,"percentLowAir":38.548063,"percentHighAir":2.03011925},"positioning":{"avgDistanceToBall":11848,"avgDistanceToBallPossession":11357,"avgDistanceToBallNoPossession":12202,"avgDistanceToMates":15030,"timeDefensiveThird":668.48,"timeNeutralThird":396.78999999999996,"timeOffensiveThird":267.03,"timeDefensiveHalf":884.37,"timeOffensiveHalf":447.85999999999996,"timeBehindBall":929.46,"timeInfrontBall":402.87,"timeMostBack":438.1,"timeMostForward":418.9,"goalsAgainstWhileLastDefender":3,"timeClosestToBall":443.1,"timeFarthestFromBall":400.3,"percentDefensiveThird":49.940807750000005,"percentOffensiveThird":20.1722825,"percentNeutralThird":29.886910500000003,"percentDefensiveHalf":66.17408025,"percentOffensiveHalf":33.825916,"percentBehindBall":69.61534325,"percentInfrontBall":30.384655249999998,"percentMostBack":33.733705,"percentMostForward":32.331874,"percentClosestToBall":34.177144250000005,"percentFarthestFromBall":30.8951615},"demo":{"inflicted":2,"taken":2}},"advanced":{"goalParticipation":20,"rating":0.48669378701679367}},{"player":{"_id":"5f3d8fdd95f40596eae23d7b","slug":"3d7b-jacob","tag":"Jacob","country":"us"},"stats":{"core":{"shots":7,"goals":3,"saves":7,"assists":1,"score":945,"shootingPercentage":42.857142857142854},"boost":{"bpm":1838,"bcpm":1909.13261,"avgAmount":196.67999999999998,"amountCollected":10299,"amountStolen":2810,"amountCollectedBig":7639,"amountStolenBig":2112,"amountCollectedSmall":2660,"amountStolenSmall":698,"countCollectedBig":88,"countStolenBig":24,"countCollectedSmall":227,"countStolenSmall":55,"amountOverfill":1132,"amountOverfillStolen":287,"amountUsedWhileSupersonic":1811,"timeZeroBoost":232.37,"percentZeroBoost":18.15785525,"timeFullBoost":166.25,"percentFullBoost":12.861310750000001,"timeBoost0To25":476.43000000000006,"timeBoost25To50":294.25,"timeBoost50To75":206.85,"timeBoost75To100":330.2,"percentBoost0To25":36.565467,"percentBoost25To50":22.37562675,"percentBoost50To75":15.784209,"percentBoost75To100":25.2746945},"movement":{"avgSpeed":6587,"totalDistance":2099739,"timeSupersonicSpeed":213.03,"timeBoostSpeed":665.11,"timeSlowSpeed":451.78,"timeGround":753.0300000000001,"timeLowAir":506.35,"timeHighAir":70.56,"timePowerslide":30.5,"countPowerslide":212,"avgPowerslideDuration":0.5700000000000001,"avgSpeedPercentage":71.597829,"percentSlowSpeed":33.826806499999996,"percentBoostSpeed":50.110704,"percentSupersonicSpeed":16.06248875,"percentGround":56.51690525,"percentLowAir":38.201848749999996,"percentHighAir":5.281246025000001},"positioning":{"avgDistanceToBall":11762,"avgDistanceToBallPossession":10592,"avgDistanceToBallNoPossession":12764,"avgDistanceToMates":15153,"timeDefensiveThird":628.55,"timeNeutralThird":394.96,"timeOffensiveThird":306.43,"timeDefensiveHalf":838.37,"timeOffensiveHalf":491.33000000000004,"timeBehindBall":839.78,"timeInfrontBall":490.15999999999997,"timeMostBack":332.09999999999997,"timeMostForward":512.9,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":457,"timeFarthestFromBall":402.6,"percentDefensiveThird":47.13845625,"percentOffensiveThird":23.20587325,"percentNeutralThird":29.65567175,"percentDefensiveHalf":62.8753545,"percentOffensiveHalf":37.12464325,"percentBehindBall":62.932387500000004,"percentInfrontBall":37.067608750000005,"percentMostBack":25.67395625,"percentMostForward":39.573733000000004,"percentClosestToBall":35.323877499999995,"percentFarthestFromBall":31.031751},"demo":{"inflicted":5,"taken":3}},"advanced":{"goalParticipation":80,"rating":0.7167417224429036}},{"player":{"_id":"5f3d8fdd95f40596eae23d7c","slug":"3d7c-sadjunior","tag":"Sadjunior","country":"ca"},"stats":{"core":{"shots":7,"goals":1,"saves":6,"assists":2,"score":820,"shootingPercentage":14.285714285714285},"boost":{"bpm":1513,"bcpm":1524.35583,"avgAmount":176.99,"amountCollected":8236,"amountStolen":1426,"amountCollectedBig":5587,"amountStolenBig":960,"amountCollectedSmall":2649,"amountStolenSmall":466,"countCollectedBig":64,"countStolenBig":10,"countCollectedSmall":219,"countStolenSmall":42,"amountOverfill":752,"amountOverfillStolen":40,"amountUsedWhileSupersonic":1564,"timeZeroBoost":250.79999999999998,"percentZeroBoost":19.3204715,"timeFullBoost":109.66,"percentFullBoost":8.4527125,"timeBoost0To25":547.21,"timeBoost25To50":246.85999999999999,"timeBoost50To75":191.78000000000003,"timeBoost75To100":298.88,"percentBoost0To25":42.595499,"percentBoost25To50":19.24515825,"percentBoost50To75":14.961071,"percentBoost75To100":23.1982685},"movement":{"avgSpeed":5899,"totalDistance":1885846,"timeSupersonicSpeed":170.33,"timeBoostSpeed":525.36,"timeSlowSpeed":639,"timeGround":806.1,"timeLowAir":474.78999999999996,"timeHighAir":53.79,"timePowerslide":21.97,"countPowerslide":174,"avgPowerslideDuration":0.51,"avgSpeedPercentage":64.119566,"percentSlowSpeed":47.83721249999999,"percentBoostSpeed":39.39101425,"percentSupersonicSpeed":12.771769125,"percentGround":60.52723475,"percentLowAir":35.5181625,"percentHighAir":3.95460395},"positioning":{"avgDistanceToBall":12870,"avgDistanceToBallPossession":12760,"avgDistanceToBallNoPossession":13212,"avgDistanceToMates":15410,"timeDefensiveThird":735.71,"timeNeutralThird":391.96,"timeOffensiveThird":207,"timeDefensiveHalf":954.05,"timeOffensiveHalf":380.34000000000003,"timeBehindBall":972.65,"timeInfrontBall":362.03,"timeMostBack":564.7,"timeMostForward":374.29999999999995,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":406.79999999999995,"timeFarthestFromBall":524.9000000000001,"percentDefensiveThird":55.00102,"percentOffensiveThird":15.675529749999999,"percentNeutralThird":29.3234465,"percentDefensiveHalf":71.34214324999999,"percentOffensiveHalf":28.657856499999998,"percentBehindBall":72.660555,"percentInfrontBall":27.339448750000003,"percentMostBack":43.66288874999999,"percentMostForward":28.904297,"percentClosestToBall":31.371422499999998,"percentFarthestFromBall":40.58687499999999},"demo":{"inflicted":1,"taken":1}},"advanced":{"goalParticipation":60,"rating":0.584025439789442}}]},"number":9,"games":[{"_id":"6082fb4c0d9dcf9da5a4d2ea","blue":1,"orange":5,"duration":300,"ballchasing":"dd308bf5-5678-4e79-96ea-c937dc631b41"},{"_id":"6082fb4d0d9dcf9da5a4d2f1","blue":1,"orange":0,"duration":300,"ballchasing":"145598b6-3116-40f5-a74f-7fa0b5b3dabf"},{"_id":"6082fb4d0d9dcf9da5a4d2f8","blue":3,"orange":0,"duration":300,"ballchasing":"d409cce2-e03d-4f5b-9e8f-63507e976adb"},{"_id":"6082fb4e0d9dcf9da5a4d2ff","blue":1,"orange":0,"duration":300,"ballchasing":"4bec9148-fbac-4b93-b500-6730f7ae833a"}]}]}

## Event [/events/{id}/participants]

- Parameters
  + id: 5f35882d53fbbb5894b43040 (required, string) - an event id

### Get Event Participants [GET]

- Response 200 (application/json)
  {"participants":[{"team":{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"},"players":[{"_id":"5f3d8fdd95f40596eae23d98","slug":"3d98-markydooda","tag":"Markydooda","country":"ab"},{"_id":"6082fb1d0d9dcf9da5a4d079","slug":"d079-greazymeister","tag":"gReazymeister"},{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","country":"it"}]},{"team":{"_id":"6020bc70f1e4807cc70023a9","slug":"23a9-genesis","name":"Genesis","image":"https://griffon.octane.gg/teams/genesis.png"},"players":[{"_id":"5f3d8fdd95f40596eae23d74","slug":"3d74-espeon","tag":"Espeon","country":"us"},{"_id":"5f3d8fdd95f40596eae23d6b","slug":"3d6b-klassux","tag":"Klassux","country":"us"},{"_id":"5f3d8fdd95f40596eae23d76","slug":"3d76-pluto","tag":"Pluto","country":"us"}]},{"team":{"_id":"6020bccdf1e4807cc70062b6","slug":"62b6-mock-it-aces","name":"Mock-It Aces","image":"https://griffon.octane.gg/teams/mock-it-aces.png"},"players":[{"_id":"5f3d8fdd95f40596eae23da0","slug":"3da0-deevo","tag":"Deevo","country":"en"},{"_id":"5f3d8fdd95f40596eae23da4","slug":"3da4-paschy90","tag":"Paschy90","country":"de"},{"_id":"5f3d8fdd95f40596eae23d9c","slug":"3d9c-violentpanda","tag":"ViolentPanda","country":"nl"}]},{"team":{"_id":"6020bc70f1e4807cc70023a0","slug":"23a0-nrg-esports","name":"NRG Esports","image":"https://griffon.octane.gg/teams/nrg-esports.png"},"players":[{"_id":"5f3d8fdd95f40596eae23d7a","slug":"3d7a-fireburner","tag":"Fireburner","country":"us"},{"_id":"5f3d8fdd95f40596eae23d7b","slug":"3d7b-jacob","tag":"Jacob","country":"us"},{"_id":"5f3d8fdd95f40596eae23d7c","slug":"3d7c-sadjunior","tag":"Sadjunior","country":"ca"}]},{"team":{"_id":"6020bccdf1e4807cc70062de","slug":"62de-northern-gaming","name":"Northern Gaming","image":"https://griffon.octane.gg/teams/northern-gaming.png"},"players":[{"_id":"5f3d8fdd95f40596eae23da6","slug":"3da6-maestro","tag":"Maestro","country":"dk"},{"_id":"5f3d8fdd95f40596eae23d99","slug":"3d99-miztik","tag":"miztik","country":"ab"},{"_id":"5f9c7dc65246bf27936b664b","slug":"664b-remkoe","tag":"remkoe","country":"nl"}]},{"team":{"_id":"6020bc70f1e4807cc7002413","slug":"2413-orbit","name":"Orbit","image":"https://griffon.octane.gg/teams/orbit.png"},"players":[{"_id":"5f3d8fdd95f40596eae23d6f","slug":"3d6f-garrettg","tag":"GarrettG","country":"us"},{"_id":"5f3d8fdd95f40596eae23d70","slug":"3d70-moses","tag":"Moses","country":"us"},{"_id":"5f3d8fdd95f40596eae23d73","slug":"3d73-turtle","tag":"Turtle","country":"us"}]},{"team":{"_id":"6020bc70f1e4807cc70023af","slug":"23af-precision-z","name":"Precision Z","image":"https://griffon.octane.gg/teams/precision-z.png"},"players":[{"_id":"5f3d8fdd95f40596eae23d9a","slug":"3d9a-kaydop","tag":"Kaydop","country":"fr"},{"_id":"5f3d8fdd95f40596eae23d9e","slug":"3d9e-sikii","tag":"Sikii","country":"de"},{"_id":"5f3d8fdd95f40596eae23e77","slug":"3e77-skyline","tag":"Skyline","country":"ch"}]},{"team":{"_id":"6020bc70f1e4807cc700242f","slug":"242f-take-3","name":"Take 3","image":"https://griffon.octane.gg/teams/take-3.png"},"players":[{"_id":"5f3d8fdd95f40596eae23d75","slug":"3d75-insolences","tag":"Insolences","country":"us"},{"_id":"5f3d8fdd95f40596eae23d78","slug":"3d78-rizzo","tag":"Rizzo","country":"us"},{"_id":"5f3d8fdd95f40596eae23d79","slug":"3d79-zanejackey","tag":"Zanejackey","country":"us"}]}]}

# Group Matches

## Matches [/matches]

### List Matches [GET]

+ Parameters
    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + qualifier: true (boolean, optional) - stage is a qualifier
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
      
    + reverseSweep: true (boolean, optional) - match is a reverse sweep

    + reverseSweepAttempt: true (boolean, optional) - match is a reverse sweep attempt
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id
    
    + tbd: false (boolean, optional) - no tbd matches
    
    + sort: name (string, optional) - field to sort by
    
    + order: `asc` (enum[string], optional) - order of sort
        + Members
          + asc
          + desc
    
    + page: 1 (int) - page number
        + Default: 1
    
    + perPage: 20 (int) - results per page
        + Default: 50

* Response 200 (application/json)
  {"matches":[{"_id":"6043152fa09e7fba40d2ae62","slug":"ae62-flipsid3-tactics-vs-nrg-esports","octane_id":"0350109","event":{"_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","region":"INT","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"]},"stage":{"_id":0,"name":"Main Event","format":"bracket-8de"},"date":"2016-12-04T00:00:00Z","format":{"type":"best","length":5},"blue":{"score":3,"winner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"},"stats":{"core":{"shots":34,"goals":6,"saves":14,"assists":5,"score":2620,"shootingPercentage":17.647058823529413},"boost":{"bpm":4841,"bcpm":4951.03124,"avgAmount":600.61,"amountCollected":26736,"amountStolen":5321,"amountCollectedBig":19445,"amountStolenBig":3178,"amountCollectedSmall":7291,"amountStolenSmall":2143,"countCollectedBig":225,"countStolenBig":39,"countCollectedSmall":653,"countStolenSmall":184,"amountOverfill":3162,"amountOverfillStolen":647,"amountUsedWhileSupersonic":3745,"timeZeroBoost":420.44,"timeFullBoost":418.98,"timeBoost0To25":1358.4,"timeBoost25To50":923.6800000000001,"timeBoost50To75":754.07,"timeBoost75To100":891.18},"movement":{"totalDistance":5828633,"timeSupersonicSpeed":551.12,"timeBoostSpeed":1636.22,"timeSlowSpeed":1803.4099999999999,"timeGround":2299.33,"timeLowAir":1528.4899999999998,"timeHighAir":162.89999999999998,"timePowerslide":90.80999999999999,"countPowerslide":770},"positioning":{"timeDefensiveThird":1820.74,"timeNeutralThird":1319.62,"timeOffensiveThird":850.3599999999999,"timeDefensiveHalf":2525.01,"timeOffensiveHalf":1465.1799999999998,"timeBehindBall":2979.85,"timeInfrontBall":1010.86},"demo":{"inflicted":6,"taken":8}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","country":"it"},"stats":{"core":{"shots":9,"goals":3,"saves":5,"assists":2,"score":940,"shootingPercentage":33.33333333333333},"boost":{"bpm":1355,"bcpm":1383.3523500000001,"avgAmount":209.42000000000002,"amountCollected":7488,"amountStolen":1478,"amountCollectedBig":4925,"amountStolenBig":814,"amountCollectedSmall":2563,"amountStolenSmall":664,"countCollectedBig":60,"countStolenBig":11,"countCollectedSmall":225,"countStolenSmall":56,"amountOverfill":1048,"amountOverfillStolen":196,"amountUsedWhileSupersonic":889,"timeZeroBoost":73.11,"percentZeroBoost":5.66867,"timeFullBoost":103.88999999999999,"percentFullBoost":7.945092724999999,"timeBoost0To25":371.58000000000004,"timeBoost25To50":339.95000000000005,"timeBoost50To75":302.69,"timeBoost75To100":293.02,"percentBoost0To25":28.43400725,"percentBoost25To50":26.14506925,"percentBoost50To75":23.053028750000003,"percentBoost75To100":22.3678945},"movement":{"avgSpeed":5746,"totalDistance":1836792,"timeSupersonicSpeed":132.97,"timeBoostSpeed":522.86,"timeSlowSpeed":677.12,"timeGround":769.29,"timeLowAir":515.0999999999999,"timeHighAir":48.55,"timePowerslide":29.229999999999997,"countPowerslide":228,"avgPowerslideDuration":0.51,"avgSpeedPercentage":62.456519,"percentSlowSpeed":50.90838250000001,"percentBoostSpeed":39.1156835,"percentSupersonicSpeed":9.97593335,"percentGround":57.59838475,"percentLowAir":38.7270215,"percentHighAir":3.67459345},"positioning":{"avgDistanceToBall":12416,"avgDistanceToBallPossession":12348,"avgDistanceToBallNoPossession":12606,"avgDistanceToMates":14960,"timeDefensiveThird":644.42,"timeNeutralThird":447.56,"timeOffensiveThird":240.95999999999998,"timeDefensiveHalf":893.57,"timeOffensiveHalf":439.23,"timeBehindBall":1049.42,"timeInfrontBall":283.51000000000005,"timeMostBack":518,"timeMostForward":331.6,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":418.9,"timeFarthestFromBall":425.80000000000007,"percentDefensiveThird":48.467209999999994,"percentOffensiveThird":18.10444525,"percentNeutralThird":33.42834875,"percentDefensiveHalf":67.11758375,"percentOffensiveHalf":32.882416,"percentBehindBall":78.83672,"percentInfrontBall":21.16328125,"percentMostBack":40.0437585,"percentMostForward":25.44998075,"percentClosestToBall":32.311219,"percentFarthestFromBall":32.8769945},"demo":{"inflicted":4,"taken":2}},"advanced":{"goalParticipation":83.33333333333334,"rating":1.0904344687742749}},{"player":{"_id":"6082fb1d0d9dcf9da5a4d079","slug":"d079-greazymeister","tag":"gReazymeister"},"stats":{"core":{"shots":11,"goals":1,"saves":6,"assists":1,"score":855,"shootingPercentage":9.090909090909092},"boost":{"bpm":1749,"bcpm":1780.08505,"avgAmount":191.92000000000002,"amountCollected":9610,"amountStolen":1666,"amountCollectedBig":7141,"amountStolenBig":951,"amountCollectedSmall":2469,"amountStolenSmall":715,"countCollectedBig":80,"countStolenBig":11,"countCollectedSmall":224,"countStolenSmall":59,"amountOverfill":925,"amountOverfillStolen":145,"amountUsedWhileSupersonic":1232,"timeZeroBoost":193.96,"percentZeroBoost":15.053702249999999,"timeFullBoost":118.42999999999999,"percentFullBoost":9.116349875000001,"timeBoost0To25":543.5799999999999,"timeBoost25To50":265.74,"timeBoost50To75":222.40000000000003,"timeBoost75To100":286.5,"percentBoost0To25":41.26221275,"percentBoost25To50":20.05843075,"percentBoost50To75":16.95397875,"percentBoost75To100":21.72537775},"movement":{"avgSpeed":6211,"totalDistance":1977383,"timeSupersonicSpeed":191.99,"timeBoostSpeed":568.89,"timeSlowSpeed":567.1500000000001,"timeGround":788.93,"timeLowAir":481.87,"timeHighAir":57.209999999999994,"timePowerslide":30.569999999999993,"countPowerslide":235,"avgPowerslideDuration":0.52,"avgSpeedPercentage":67.51087150000001,"percentSlowSpeed":42.834596499999996,"percentBoostSpeed":42.83116575,"percentSupersonicSpeed":14.33424075,"percentGround":59.40004425,"percentLowAir":36.317852,"percentHighAir":4.282102275},"positioning":{"avgDistanceToBall":12676,"avgDistanceToBallPossession":12163,"avgDistanceToBallNoPossession":13231,"avgDistanceToMates":15338,"timeDefensiveThird":633.36,"timeNeutralThird":431.29,"timeOffensiveThird":263.37,"timeDefensiveHalf":869.01,"timeOffensiveHalf":458.75,"timeBehindBall":1001.28,"timeInfrontBall":326.74,"timeMostBack":480.20000000000005,"timeMostForward":452.7,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":449.4,"timeFarthestFromBall":436.9,"percentDefensiveThird":47.78079675,"percentOffensiveThird":19.832879125,"percentNeutralThird":32.386322,"percentDefensiveHalf":65.47910325,"percentOffensiveHalf":34.520901,"percentBehindBall":75.44333999999999,"percentInfrontBall":24.556661000000002,"percentMostBack":37.11360125,"percentMostForward":34.8868515,"percentClosestToBall":34.765083749999995,"percentFarthestFromBall":33.646304},"demo":{"inflicted":2,"taken":3}},"advanced":{"goalParticipation":33.33333333333333,"rating":0.6576622683574388}},{"player":{"_id":"5f3d8fdd95f40596eae23d98","slug":"3d98-markydooda","tag":"Markydooda","country":"ab"},"stats":{"core":{"shots":14,"goals":2,"saves":3,"assists":2,"score":825,"shootingPercentage":14.285714285714285},"boost":{"bpm":1737,"bcpm":1787.59384,"avgAmount":199.27,"amountCollected":9638,"amountStolen":2177,"amountCollectedBig":7379,"amountStolenBig":1413,"amountCollectedSmall":2259,"amountStolenSmall":764,"countCollectedBig":85,"countStolenBig":17,"countCollectedSmall":204,"countStolenSmall":69,"amountOverfill":1189,"amountOverfillStolen":306,"amountUsedWhileSupersonic":1624,"timeZeroBoost":153.37,"percentZeroBoost":11.839296999999998,"timeFullBoost":196.66000000000003,"percentFullBoost":15.160770750000001,"timeBoost0To25":443.24,"timeBoost25To50":317.99,"timeBoost50To75":228.98000000000002,"timeBoost75To100":311.65999999999997,"percentBoost0To25":34.12572475,"percentBoost25To50":24.450708749999997,"percentBoost50To75":17.511339625,"percentBoost75To100":23.91222475},"movement":{"avgSpeed":6325,"totalDistance":2014458,"timeSupersonicSpeed":226.16,"timeBoostSpeed":544.47,"timeSlowSpeed":559.14,"timeGround":741.1099999999999,"timeLowAir":531.52,"timeHighAir":57.14,"timePowerslide":31.01,"countPowerslide":307,"avgPowerslideDuration":0.4,"avgSpeedPercentage":68.7500005,"percentSlowSpeed":41.99217525,"percentBoostSpeed":40.93057,"percentSupersonicSpeed":17.0772575,"percentGround":55.656828749999995,"percentLowAir":40.04468875,"percentHighAir":4.298480475},"positioning":{"avgDistanceToBall":12663,"avgDistanceToBallPossession":12355,"avgDistanceToBallNoPossession":12989,"avgDistanceToMates":15613,"timeDefensiveThird":542.96,"timeNeutralThird":440.77,"timeOffensiveThird":346.03,"timeDefensiveHalf":762.4300000000001,"timeOffensiveHalf":567.1999999999999,"timeBehindBall":929.1500000000001,"timeInfrontBall":400.61,"timeMostBack":331.1,"timeMostForward":523.9,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":440.30000000000007,"timeFarthestFromBall":466.5,"percentDefensiveThird":40.66040425,"percentOffensiveThird":26.109765250000002,"percentNeutralThird":33.2298305,"percentDefensiveHalf":57.213438999999994,"percentOffensiveHalf":42.78656325,"percentBehindBall":69.843096,"percentInfrontBall":30.156903500000002,"percentMostBack":25.40523275,"percentMostForward":40.625619,"percentClosestToBall":33.915815,"percentFarthestFromBall":36.031493499999996},"demo":{"inflicted":0,"taken":3}},"advanced":{"goalParticipation":66.66666666666666,"rating":0.9047666428357998}}]},"orange":{"score":1,"team":{"team":{"_id":"6020bc70f1e4807cc70023a0","slug":"23a0-nrg-esports","name":"NRG Esports","image":"https://griffon.octane.gg/teams/nrg-esports.png"},"stats":{"core":{"shots":22,"goals":5,"saves":22,"assists":3,"score":2540,"shootingPercentage":22.727272727272727},"boost":{"bpm":5095,"bcpm":5181.857169999999,"avgAmount":568.9,"amountCollected":27984,"amountStolen":6409,"amountCollectedBig":20642,"amountStolenBig":4724,"amountCollectedSmall":7342,"amountStolenSmall":1685,"countCollectedBig":237,"countStolenBig":53,"countCollectedSmall":632,"countStolenSmall":148,"amountOverfill":2954,"amountOverfillStolen":547,"amountUsedWhileSupersonic":5117,"timeZeroBoost":639.2099999999999,"timeFullBoost":445.38,"timeBoost0To25":1479.33,"timeBoost25To50":829.08,"timeBoost50To75":667.6700000000001,"timeBoost75To100":919.11},"movement":{"totalDistance":5923761,"timeSupersonicSpeed":595.7900000000001,"timeBoostSpeed":1719.2600000000002,"timeSlowSpeed":1681.8899999999999,"timeGround":2351.83,"timeLowAir":1494.1999999999998,"timeHighAir":150.91,"timePowerslide":73.97999999999999,"countPowerslide":550},"positioning":{"timeDefensiveThird":2032.74,"timeNeutralThird":1183.71,"timeOffensiveThird":780.46,"timeDefensiveHalf":2676.79,"timeOffensiveHalf":1319.5300000000002,"timeBehindBall":2741.89,"timeInfrontBall":1255.06},"demo":{"inflicted":8,"taken":6}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7a","slug":"3d7a-fireburner","tag":"Fireburner","country":"us"},"stats":{"core":{"shots":8,"goals":1,"saves":9,"assists":0,"score":775,"shootingPercentage":12.5},"boost":{"bpm":1744,"bcpm":1748.36873,"avgAmount":195.23,"amountCollected":9449,"amountStolen":2173,"amountCollectedBig":7416,"amountStolenBig":1652,"amountCollectedSmall":2033,"amountStolenSmall":521,"countCollectedBig":85,"countStolenBig":19,"countCollectedSmall":186,"countStolenSmall":51,"amountOverfill":1070,"amountOverfillStolen":220,"amountUsedWhileSupersonic":1742,"timeZeroBoost":156.04,"percentZeroBoost":12.079158625000002,"timeFullBoost":169.47,"percentFullBoost":12.9679785,"timeBoost0To25":455.68999999999994,"timeBoost25To50":287.96999999999997,"timeBoost50To75":269.03999999999996,"timeBoost75To100":290.03,"percentBoost0To25":34.95103375,"percentBoost25To50":22.17023925,"percentBoost50To75":20.63573975,"percentBoost75To100":22.24298725},"movement":{"avgSpeed":6069,"totalDistance":1938176,"timeSupersonicSpeed":212.43,"timeBoostSpeed":528.79,"timeSlowSpeed":591.11,"timeGround":792.7,"timeLowAir":513.06,"timeHighAir":26.560000000000002,"timePowerslide":21.509999999999998,"countPowerslide":164,"avgPowerslideDuration":0.52,"avgSpeedPercentage":65.967393,"percentSlowSpeed":44.25205375,"percentBoostSpeed":39.8215655,"percentSupersonicSpeed":15.926380625,"percentGround":59.421816500000006,"percentLowAir":38.548063,"percentHighAir":2.03011925},"positioning":{"avgDistanceToBall":11848,"avgDistanceToBallPossession":11357,"avgDistanceToBallNoPossession":12202,"avgDistanceToMates":15030,"timeDefensiveThird":668.48,"timeNeutralThird":396.78999999999996,"timeOffensiveThird":267.03,"timeDefensiveHalf":884.37,"timeOffensiveHalf":447.85999999999996,"timeBehindBall":929.46,"timeInfrontBall":402.87,"timeMostBack":438.1,"timeMostForward":418.9,"goalsAgainstWhileLastDefender":3,"timeClosestToBall":443.1,"timeFarthestFromBall":400.3,"percentDefensiveThird":49.940807750000005,"percentOffensiveThird":20.1722825,"percentNeutralThird":29.886910500000003,"percentDefensiveHalf":66.17408025,"percentOffensiveHalf":33.825916,"percentBehindBall":69.61534325,"percentInfrontBall":30.384655249999998,"percentMostBack":33.733705,"percentMostForward":32.331874,"percentClosestToBall":34.177144250000005,"percentFarthestFromBall":30.8951615},"demo":{"inflicted":2,"taken":2}},"advanced":{"goalParticipation":20,"rating":0.48669378701679367}},{"player":{"_id":"5f3d8fdd95f40596eae23d7b","slug":"3d7b-jacob","tag":"Jacob","country":"us"},"stats":{"core":{"shots":7,"goals":3,"saves":7,"assists":1,"score":945,"shootingPercentage":42.857142857142854},"boost":{"bpm":1838,"bcpm":1909.13261,"avgAmount":196.67999999999998,"amountCollected":10299,"amountStolen":2810,"amountCollectedBig":7639,"amountStolenBig":2112,"amountCollectedSmall":2660,"amountStolenSmall":698,"countCollectedBig":88,"countStolenBig":24,"countCollectedSmall":227,"countStolenSmall":55,"amountOverfill":1132,"amountOverfillStolen":287,"amountUsedWhileSupersonic":1811,"timeZeroBoost":232.37,"percentZeroBoost":18.15785525,"timeFullBoost":166.25,"percentFullBoost":12.861310750000001,"timeBoost0To25":476.43000000000006,"timeBoost25To50":294.25,"timeBoost50To75":206.85,"timeBoost75To100":330.2,"percentBoost0To25":36.565467,"percentBoost25To50":22.37562675,"percentBoost50To75":15.784209,"percentBoost75To100":25.2746945},"movement":{"avgSpeed":6587,"totalDistance":2099739,"timeSupersonicSpeed":213.03,"timeBoostSpeed":665.11,"timeSlowSpeed":451.78,"timeGround":753.0300000000001,"timeLowAir":506.35,"timeHighAir":70.56,"timePowerslide":30.5,"countPowerslide":212,"avgPowerslideDuration":0.5700000000000001,"avgSpeedPercentage":71.597829,"percentSlowSpeed":33.826806499999996,"percentBoostSpeed":50.110704,"percentSupersonicSpeed":16.06248875,"percentGround":56.51690525,"percentLowAir":38.201848749999996,"percentHighAir":5.281246025000001},"positioning":{"avgDistanceToBall":11762,"avgDistanceToBallPossession":10592,"avgDistanceToBallNoPossession":12764,"avgDistanceToMates":15153,"timeDefensiveThird":628.55,"timeNeutralThird":394.96,"timeOffensiveThird":306.43,"timeDefensiveHalf":838.37,"timeOffensiveHalf":491.33000000000004,"timeBehindBall":839.78,"timeInfrontBall":490.15999999999997,"timeMostBack":332.09999999999997,"timeMostForward":512.9,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":457,"timeFarthestFromBall":402.6,"percentDefensiveThird":47.13845625,"percentOffensiveThird":23.20587325,"percentNeutralThird":29.65567175,"percentDefensiveHalf":62.8753545,"percentOffensiveHalf":37.12464325,"percentBehindBall":62.932387500000004,"percentInfrontBall":37.067608750000005,"percentMostBack":25.67395625,"percentMostForward":39.573733000000004,"percentClosestToBall":35.323877499999995,"percentFarthestFromBall":31.031751},"demo":{"inflicted":5,"taken":3}},"advanced":{"goalParticipation":80,"rating":0.7167417224429036}},{"player":{"_id":"5f3d8fdd95f40596eae23d7c","slug":"3d7c-sadjunior","tag":"Sadjunior","country":"ca"},"stats":{"core":{"shots":7,"goals":1,"saves":6,"assists":2,"score":820,"shootingPercentage":14.285714285714285},"boost":{"bpm":1513,"bcpm":1524.35583,"avgAmount":176.99,"amountCollected":8236,"amountStolen":1426,"amountCollectedBig":5587,"amountStolenBig":960,"amountCollectedSmall":2649,"amountStolenSmall":466,"countCollectedBig":64,"countStolenBig":10,"countCollectedSmall":219,"countStolenSmall":42,"amountOverfill":752,"amountOverfillStolen":40,"amountUsedWhileSupersonic":1564,"timeZeroBoost":250.79999999999998,"percentZeroBoost":19.3204715,"timeFullBoost":109.66,"percentFullBoost":8.4527125,"timeBoost0To25":547.21,"timeBoost25To50":246.85999999999999,"timeBoost50To75":191.78000000000003,"timeBoost75To100":298.88,"percentBoost0To25":42.595499,"percentBoost25To50":19.24515825,"percentBoost50To75":14.961071,"percentBoost75To100":23.1982685},"movement":{"avgSpeed":5899,"totalDistance":1885846,"timeSupersonicSpeed":170.33,"timeBoostSpeed":525.36,"timeSlowSpeed":639,"timeGround":806.1,"timeLowAir":474.78999999999996,"timeHighAir":53.79,"timePowerslide":21.97,"countPowerslide":174,"avgPowerslideDuration":0.51,"avgSpeedPercentage":64.119566,"percentSlowSpeed":47.83721249999999,"percentBoostSpeed":39.39101425,"percentSupersonicSpeed":12.771769125,"percentGround":60.52723475,"percentLowAir":35.5181625,"percentHighAir":3.95460395},"positioning":{"avgDistanceToBall":12870,"avgDistanceToBallPossession":12760,"avgDistanceToBallNoPossession":13212,"avgDistanceToMates":15410,"timeDefensiveThird":735.71,"timeNeutralThird":391.96,"timeOffensiveThird":207,"timeDefensiveHalf":954.05,"timeOffensiveHalf":380.34000000000003,"timeBehindBall":972.65,"timeInfrontBall":362.03,"timeMostBack":564.7,"timeMostForward":374.29999999999995,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":406.79999999999995,"timeFarthestFromBall":524.9000000000001,"percentDefensiveThird":55.00102,"percentOffensiveThird":15.675529749999999,"percentNeutralThird":29.3234465,"percentDefensiveHalf":71.34214324999999,"percentOffensiveHalf":28.657856499999998,"percentBehindBall":72.660555,"percentInfrontBall":27.339448750000003,"percentMostBack":43.66288874999999,"percentMostForward":28.904297,"percentClosestToBall":31.371422499999998,"percentFarthestFromBall":40.58687499999999},"demo":{"inflicted":1,"taken":1}},"advanced":{"goalParticipation":60,"rating":0.584025439789442}}]},"number":9,"games":[{"_id":"6082fb4c0d9dcf9da5a4d2ea","blue":1,"orange":5,"duration":300,"ballchasing":"dd308bf5-5678-4e79-96ea-c937dc631b41"},{"_id":"6082fb4d0d9dcf9da5a4d2f1","blue":1,"orange":0,"duration":300,"ballchasing":"145598b6-3116-40f5-a74f-7fa0b5b3dabf"},{"_id":"6082fb4d0d9dcf9da5a4d2f8","blue":3,"orange":0,"duration":300,"ballchasing":"d409cce2-e03d-4f5b-9e8f-63507e976adb"},{"_id":"6082fb4e0d9dcf9da5a4d2ff","blue":1,"orange":0,"duration":300,"ballchasing":"4bec9148-fbac-4b93-b500-6730f7ae833a"}]}]}

## Match [/matches/{id}]

- Parameters
  + id: 6043152fa09e7fba40d2ae62 (required, string) - a match id

### Get Match [GET]

- Response 200 (application/json)
{"_id":"6043152fa09e7fba40d2ae62","slug":"ae62-flipsid3-tactics-vs-nrg-esports","octane_id":"0350109","event":{"_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","region":"INT","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"]},"stage":{"_id":0,"name":"Main Event","format":"bracket-8de"},"date":"2016-12-04T00:00:00Z","format":{"type":"best","length":5},"blue":{"score":3,"winner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"},"stats":{"core":{"shots":34,"goals":6,"saves":14,"assists":5,"score":2620,"shootingPercentage":17.647058823529413},"boost":{"bpm":4841,"bcpm":4951.03124,"avgAmount":600.61,"amountCollected":26736,"amountStolen":5321,"amountCollectedBig":19445,"amountStolenBig":3178,"amountCollectedSmall":7291,"amountStolenSmall":2143,"countCollectedBig":225,"countStolenBig":39,"countCollectedSmall":653,"countStolenSmall":184,"amountOverfill":3162,"amountOverfillStolen":647,"amountUsedWhileSupersonic":3745,"timeZeroBoost":420.44,"timeFullBoost":418.98,"timeBoost0To25":1358.4,"timeBoost25To50":923.6800000000001,"timeBoost50To75":754.07,"timeBoost75To100":891.18},"movement":{"totalDistance":5828633,"timeSupersonicSpeed":551.12,"timeBoostSpeed":1636.22,"timeSlowSpeed":1803.4099999999999,"timeGround":2299.33,"timeLowAir":1528.4899999999998,"timeHighAir":162.89999999999998,"timePowerslide":90.80999999999999,"countPowerslide":770},"positioning":{"timeDefensiveThird":1820.74,"timeNeutralThird":1319.62,"timeOffensiveThird":850.3599999999999,"timeDefensiveHalf":2525.01,"timeOffensiveHalf":1465.1799999999998,"timeBehindBall":2979.85,"timeInfrontBall":1010.86},"demo":{"inflicted":6,"taken":8}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","country":"it"},"stats":{"core":{"shots":9,"goals":3,"saves":5,"assists":2,"score":940,"shootingPercentage":33.33333333333333},"boost":{"bpm":1355,"bcpm":1383.3523500000001,"avgAmount":209.42000000000002,"amountCollected":7488,"amountStolen":1478,"amountCollectedBig":4925,"amountStolenBig":814,"amountCollectedSmall":2563,"amountStolenSmall":664,"countCollectedBig":60,"countStolenBig":11,"countCollectedSmall":225,"countStolenSmall":56,"amountOverfill":1048,"amountOverfillStolen":196,"amountUsedWhileSupersonic":889,"timeZeroBoost":73.11,"percentZeroBoost":5.66867,"timeFullBoost":103.88999999999999,"percentFullBoost":7.945092724999999,"timeBoost0To25":371.58000000000004,"timeBoost25To50":339.95000000000005,"timeBoost50To75":302.69,"timeBoost75To100":293.02,"percentBoost0To25":28.43400725,"percentBoost25To50":26.14506925,"percentBoost50To75":23.053028750000003,"percentBoost75To100":22.3678945},"movement":{"avgSpeed":5746,"totalDistance":1836792,"timeSupersonicSpeed":132.97,"timeBoostSpeed":522.86,"timeSlowSpeed":677.12,"timeGround":769.29,"timeLowAir":515.0999999999999,"timeHighAir":48.55,"timePowerslide":29.229999999999997,"countPowerslide":228,"avgPowerslideDuration":0.51,"avgSpeedPercentage":62.456519,"percentSlowSpeed":50.90838250000001,"percentBoostSpeed":39.1156835,"percentSupersonicSpeed":9.97593335,"percentGround":57.59838475,"percentLowAir":38.7270215,"percentHighAir":3.67459345},"positioning":{"avgDistanceToBall":12416,"avgDistanceToBallPossession":12348,"avgDistanceToBallNoPossession":12606,"avgDistanceToMates":14960,"timeDefensiveThird":644.42,"timeNeutralThird":447.56,"timeOffensiveThird":240.95999999999998,"timeDefensiveHalf":893.57,"timeOffensiveHalf":439.23,"timeBehindBall":1049.42,"timeInfrontBall":283.51000000000005,"timeMostBack":518,"timeMostForward":331.6,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":418.9,"timeFarthestFromBall":425.80000000000007,"percentDefensiveThird":48.467209999999994,"percentOffensiveThird":18.10444525,"percentNeutralThird":33.42834875,"percentDefensiveHalf":67.11758375,"percentOffensiveHalf":32.882416,"percentBehindBall":78.83672,"percentInfrontBall":21.16328125,"percentMostBack":40.0437585,"percentMostForward":25.44998075,"percentClosestToBall":32.311219,"percentFarthestFromBall":32.8769945},"demo":{"inflicted":4,"taken":2}},"advanced":{"goalParticipation":83.33333333333334,"rating":1.0904344687742749}},{"player":{"_id":"6082fb1d0d9dcf9da5a4d079","slug":"d079-greazymeister","tag":"gReazymeister"},"stats":{"core":{"shots":11,"goals":1,"saves":6,"assists":1,"score":855,"shootingPercentage":9.090909090909092},"boost":{"bpm":1749,"bcpm":1780.08505,"avgAmount":191.92000000000002,"amountCollected":9610,"amountStolen":1666,"amountCollectedBig":7141,"amountStolenBig":951,"amountCollectedSmall":2469,"amountStolenSmall":715,"countCollectedBig":80,"countStolenBig":11,"countCollectedSmall":224,"countStolenSmall":59,"amountOverfill":925,"amountOverfillStolen":145,"amountUsedWhileSupersonic":1232,"timeZeroBoost":193.96,"percentZeroBoost":15.053702249999999,"timeFullBoost":118.42999999999999,"percentFullBoost":9.116349875000001,"timeBoost0To25":543.5799999999999,"timeBoost25To50":265.74,"timeBoost50To75":222.40000000000003,"timeBoost75To100":286.5,"percentBoost0To25":41.26221275,"percentBoost25To50":20.05843075,"percentBoost50To75":16.95397875,"percentBoost75To100":21.72537775},"movement":{"avgSpeed":6211,"totalDistance":1977383,"timeSupersonicSpeed":191.99,"timeBoostSpeed":568.89,"timeSlowSpeed":567.1500000000001,"timeGround":788.93,"timeLowAir":481.87,"timeHighAir":57.209999999999994,"timePowerslide":30.569999999999993,"countPowerslide":235,"avgPowerslideDuration":0.52,"avgSpeedPercentage":67.51087150000001,"percentSlowSpeed":42.834596499999996,"percentBoostSpeed":42.83116575,"percentSupersonicSpeed":14.33424075,"percentGround":59.40004425,"percentLowAir":36.317852,"percentHighAir":4.282102275},"positioning":{"avgDistanceToBall":12676,"avgDistanceToBallPossession":12163,"avgDistanceToBallNoPossession":13231,"avgDistanceToMates":15338,"timeDefensiveThird":633.36,"timeNeutralThird":431.29,"timeOffensiveThird":263.37,"timeDefensiveHalf":869.01,"timeOffensiveHalf":458.75,"timeBehindBall":1001.28,"timeInfrontBall":326.74,"timeMostBack":480.20000000000005,"timeMostForward":452.7,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":449.4,"timeFarthestFromBall":436.9,"percentDefensiveThird":47.78079675,"percentOffensiveThird":19.832879125,"percentNeutralThird":32.386322,"percentDefensiveHalf":65.47910325,"percentOffensiveHalf":34.520901,"percentBehindBall":75.44333999999999,"percentInfrontBall":24.556661000000002,"percentMostBack":37.11360125,"percentMostForward":34.8868515,"percentClosestToBall":34.765083749999995,"percentFarthestFromBall":33.646304},"demo":{"inflicted":2,"taken":3}},"advanced":{"goalParticipation":33.33333333333333,"rating":0.6576622683574388}},{"player":{"_id":"5f3d8fdd95f40596eae23d98","slug":"3d98-markydooda","tag":"Markydooda","country":"ab"},"stats":{"core":{"shots":14,"goals":2,"saves":3,"assists":2,"score":825,"shootingPercentage":14.285714285714285},"boost":{"bpm":1737,"bcpm":1787.59384,"avgAmount":199.27,"amountCollected":9638,"amountStolen":2177,"amountCollectedBig":7379,"amountStolenBig":1413,"amountCollectedSmall":2259,"amountStolenSmall":764,"countCollectedBig":85,"countStolenBig":17,"countCollectedSmall":204,"countStolenSmall":69,"amountOverfill":1189,"amountOverfillStolen":306,"amountUsedWhileSupersonic":1624,"timeZeroBoost":153.37,"percentZeroBoost":11.839296999999998,"timeFullBoost":196.66000000000003,"percentFullBoost":15.160770750000001,"timeBoost0To25":443.24,"timeBoost25To50":317.99,"timeBoost50To75":228.98000000000002,"timeBoost75To100":311.65999999999997,"percentBoost0To25":34.12572475,"percentBoost25To50":24.450708749999997,"percentBoost50To75":17.511339625,"percentBoost75To100":23.91222475},"movement":{"avgSpeed":6325,"totalDistance":2014458,"timeSupersonicSpeed":226.16,"timeBoostSpeed":544.47,"timeSlowSpeed":559.14,"timeGround":741.1099999999999,"timeLowAir":531.52,"timeHighAir":57.14,"timePowerslide":31.01,"countPowerslide":307,"avgPowerslideDuration":0.4,"avgSpeedPercentage":68.7500005,"percentSlowSpeed":41.99217525,"percentBoostSpeed":40.93057,"percentSupersonicSpeed":17.0772575,"percentGround":55.656828749999995,"percentLowAir":40.04468875,"percentHighAir":4.298480475},"positioning":{"avgDistanceToBall":12663,"avgDistanceToBallPossession":12355,"avgDistanceToBallNoPossession":12989,"avgDistanceToMates":15613,"timeDefensiveThird":542.96,"timeNeutralThird":440.77,"timeOffensiveThird":346.03,"timeDefensiveHalf":762.4300000000001,"timeOffensiveHalf":567.1999999999999,"timeBehindBall":929.1500000000001,"timeInfrontBall":400.61,"timeMostBack":331.1,"timeMostForward":523.9,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":440.30000000000007,"timeFarthestFromBall":466.5,"percentDefensiveThird":40.66040425,"percentOffensiveThird":26.109765250000002,"percentNeutralThird":33.2298305,"percentDefensiveHalf":57.213438999999994,"percentOffensiveHalf":42.78656325,"percentBehindBall":69.843096,"percentInfrontBall":30.156903500000002,"percentMostBack":25.40523275,"percentMostForward":40.625619,"percentClosestToBall":33.915815,"percentFarthestFromBall":36.031493499999996},"demo":{"inflicted":0,"taken":3}},"advanced":{"goalParticipation":66.66666666666666,"rating":0.9047666428357998}}]},"orange":{"score":1,"team":{"team":{"_id":"6020bc70f1e4807cc70023a0","slug":"23a0-nrg-esports","name":"NRG Esports","image":"https://griffon.octane.gg/teams/nrg-esports.png"},"stats":{"core":{"shots":22,"goals":5,"saves":22,"assists":3,"score":2540,"shootingPercentage":22.727272727272727},"boost":{"bpm":5095,"bcpm":5181.857169999999,"avgAmount":568.9,"amountCollected":27984,"amountStolen":6409,"amountCollectedBig":20642,"amountStolenBig":4724,"amountCollectedSmall":7342,"amountStolenSmall":1685,"countCollectedBig":237,"countStolenBig":53,"countCollectedSmall":632,"countStolenSmall":148,"amountOverfill":2954,"amountOverfillStolen":547,"amountUsedWhileSupersonic":5117,"timeZeroBoost":639.2099999999999,"timeFullBoost":445.38,"timeBoost0To25":1479.33,"timeBoost25To50":829.08,"timeBoost50To75":667.6700000000001,"timeBoost75To100":919.11},"movement":{"totalDistance":5923761,"timeSupersonicSpeed":595.7900000000001,"timeBoostSpeed":1719.2600000000002,"timeSlowSpeed":1681.8899999999999,"timeGround":2351.83,"timeLowAir":1494.1999999999998,"timeHighAir":150.91,"timePowerslide":73.97999999999999,"countPowerslide":550},"positioning":{"timeDefensiveThird":2032.74,"timeNeutralThird":1183.71,"timeOffensiveThird":780.46,"timeDefensiveHalf":2676.79,"timeOffensiveHalf":1319.5300000000002,"timeBehindBall":2741.89,"timeInfrontBall":1255.06},"demo":{"inflicted":8,"taken":6}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7a","slug":"3d7a-fireburner","tag":"Fireburner","country":"us"},"stats":{"core":{"shots":8,"goals":1,"saves":9,"assists":0,"score":775,"shootingPercentage":12.5},"boost":{"bpm":1744,"bcpm":1748.36873,"avgAmount":195.23,"amountCollected":9449,"amountStolen":2173,"amountCollectedBig":7416,"amountStolenBig":1652,"amountCollectedSmall":2033,"amountStolenSmall":521,"countCollectedBig":85,"countStolenBig":19,"countCollectedSmall":186,"countStolenSmall":51,"amountOverfill":1070,"amountOverfillStolen":220,"amountUsedWhileSupersonic":1742,"timeZeroBoost":156.04,"percentZeroBoost":12.079158625000002,"timeFullBoost":169.47,"percentFullBoost":12.9679785,"timeBoost0To25":455.68999999999994,"timeBoost25To50":287.96999999999997,"timeBoost50To75":269.03999999999996,"timeBoost75To100":290.03,"percentBoost0To25":34.95103375,"percentBoost25To50":22.17023925,"percentBoost50To75":20.63573975,"percentBoost75To100":22.24298725},"movement":{"avgSpeed":6069,"totalDistance":1938176,"timeSupersonicSpeed":212.43,"timeBoostSpeed":528.79,"timeSlowSpeed":591.11,"timeGround":792.7,"timeLowAir":513.06,"timeHighAir":26.560000000000002,"timePowerslide":21.509999999999998,"countPowerslide":164,"avgPowerslideDuration":0.52,"avgSpeedPercentage":65.967393,"percentSlowSpeed":44.25205375,"percentBoostSpeed":39.8215655,"percentSupersonicSpeed":15.926380625,"percentGround":59.421816500000006,"percentLowAir":38.548063,"percentHighAir":2.03011925},"positioning":{"avgDistanceToBall":11848,"avgDistanceToBallPossession":11357,"avgDistanceToBallNoPossession":12202,"avgDistanceToMates":15030,"timeDefensiveThird":668.48,"timeNeutralThird":396.78999999999996,"timeOffensiveThird":267.03,"timeDefensiveHalf":884.37,"timeOffensiveHalf":447.85999999999996,"timeBehindBall":929.46,"timeInfrontBall":402.87,"timeMostBack":438.1,"timeMostForward":418.9,"goalsAgainstWhileLastDefender":3,"timeClosestToBall":443.1,"timeFarthestFromBall":400.3,"percentDefensiveThird":49.940807750000005,"percentOffensiveThird":20.1722825,"percentNeutralThird":29.886910500000003,"percentDefensiveHalf":66.17408025,"percentOffensiveHalf":33.825916,"percentBehindBall":69.61534325,"percentInfrontBall":30.384655249999998,"percentMostBack":33.733705,"percentMostForward":32.331874,"percentClosestToBall":34.177144250000005,"percentFarthestFromBall":30.8951615},"demo":{"inflicted":2,"taken":2}},"advanced":{"goalParticipation":20,"rating":0.48669378701679367}},{"player":{"_id":"5f3d8fdd95f40596eae23d7b","slug":"3d7b-jacob","tag":"Jacob","country":"us"},"stats":{"core":{"shots":7,"goals":3,"saves":7,"assists":1,"score":945,"shootingPercentage":42.857142857142854},"boost":{"bpm":1838,"bcpm":1909.13261,"avgAmount":196.67999999999998,"amountCollected":10299,"amountStolen":2810,"amountCollectedBig":7639,"amountStolenBig":2112,"amountCollectedSmall":2660,"amountStolenSmall":698,"countCollectedBig":88,"countStolenBig":24,"countCollectedSmall":227,"countStolenSmall":55,"amountOverfill":1132,"amountOverfillStolen":287,"amountUsedWhileSupersonic":1811,"timeZeroBoost":232.37,"percentZeroBoost":18.15785525,"timeFullBoost":166.25,"percentFullBoost":12.861310750000001,"timeBoost0To25":476.43000000000006,"timeBoost25To50":294.25,"timeBoost50To75":206.85,"timeBoost75To100":330.2,"percentBoost0To25":36.565467,"percentBoost25To50":22.37562675,"percentBoost50To75":15.784209,"percentBoost75To100":25.2746945},"movement":{"avgSpeed":6587,"totalDistance":2099739,"timeSupersonicSpeed":213.03,"timeBoostSpeed":665.11,"timeSlowSpeed":451.78,"timeGround":753.0300000000001,"timeLowAir":506.35,"timeHighAir":70.56,"timePowerslide":30.5,"countPowerslide":212,"avgPowerslideDuration":0.5700000000000001,"avgSpeedPercentage":71.597829,"percentSlowSpeed":33.826806499999996,"percentBoostSpeed":50.110704,"percentSupersonicSpeed":16.06248875,"percentGround":56.51690525,"percentLowAir":38.201848749999996,"percentHighAir":5.281246025000001},"positioning":{"avgDistanceToBall":11762,"avgDistanceToBallPossession":10592,"avgDistanceToBallNoPossession":12764,"avgDistanceToMates":15153,"timeDefensiveThird":628.55,"timeNeutralThird":394.96,"timeOffensiveThird":306.43,"timeDefensiveHalf":838.37,"timeOffensiveHalf":491.33000000000004,"timeBehindBall":839.78,"timeInfrontBall":490.15999999999997,"timeMostBack":332.09999999999997,"timeMostForward":512.9,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":457,"timeFarthestFromBall":402.6,"percentDefensiveThird":47.13845625,"percentOffensiveThird":23.20587325,"percentNeutralThird":29.65567175,"percentDefensiveHalf":62.8753545,"percentOffensiveHalf":37.12464325,"percentBehindBall":62.932387500000004,"percentInfrontBall":37.067608750000005,"percentMostBack":25.67395625,"percentMostForward":39.573733000000004,"percentClosestToBall":35.323877499999995,"percentFarthestFromBall":31.031751},"demo":{"inflicted":5,"taken":3}},"advanced":{"goalParticipation":80,"rating":0.7167417224429036}},{"player":{"_id":"5f3d8fdd95f40596eae23d7c","slug":"3d7c-sadjunior","tag":"Sadjunior","country":"ca"},"stats":{"core":{"shots":7,"goals":1,"saves":6,"assists":2,"score":820,"shootingPercentage":14.285714285714285},"boost":{"bpm":1513,"bcpm":1524.35583,"avgAmount":176.99,"amountCollected":8236,"amountStolen":1426,"amountCollectedBig":5587,"amountStolenBig":960,"amountCollectedSmall":2649,"amountStolenSmall":466,"countCollectedBig":64,"countStolenBig":10,"countCollectedSmall":219,"countStolenSmall":42,"amountOverfill":752,"amountOverfillStolen":40,"amountUsedWhileSupersonic":1564,"timeZeroBoost":250.79999999999998,"percentZeroBoost":19.3204715,"timeFullBoost":109.66,"percentFullBoost":8.4527125,"timeBoost0To25":547.21,"timeBoost25To50":246.85999999999999,"timeBoost50To75":191.78000000000003,"timeBoost75To100":298.88,"percentBoost0To25":42.595499,"percentBoost25To50":19.24515825,"percentBoost50To75":14.961071,"percentBoost75To100":23.1982685},"movement":{"avgSpeed":5899,"totalDistance":1885846,"timeSupersonicSpeed":170.33,"timeBoostSpeed":525.36,"timeSlowSpeed":639,"timeGround":806.1,"timeLowAir":474.78999999999996,"timeHighAir":53.79,"timePowerslide":21.97,"countPowerslide":174,"avgPowerslideDuration":0.51,"avgSpeedPercentage":64.119566,"percentSlowSpeed":47.83721249999999,"percentBoostSpeed":39.39101425,"percentSupersonicSpeed":12.771769125,"percentGround":60.52723475,"percentLowAir":35.5181625,"percentHighAir":3.95460395},"positioning":{"avgDistanceToBall":12870,"avgDistanceToBallPossession":12760,"avgDistanceToBallNoPossession":13212,"avgDistanceToMates":15410,"timeDefensiveThird":735.71,"timeNeutralThird":391.96,"timeOffensiveThird":207,"timeDefensiveHalf":954.05,"timeOffensiveHalf":380.34000000000003,"timeBehindBall":972.65,"timeInfrontBall":362.03,"timeMostBack":564.7,"timeMostForward":374.29999999999995,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":406.79999999999995,"timeFarthestFromBall":524.9000000000001,"percentDefensiveThird":55.00102,"percentOffensiveThird":15.675529749999999,"percentNeutralThird":29.3234465,"percentDefensiveHalf":71.34214324999999,"percentOffensiveHalf":28.657856499999998,"percentBehindBall":72.660555,"percentInfrontBall":27.339448750000003,"percentMostBack":43.66288874999999,"percentMostForward":28.904297,"percentClosestToBall":31.371422499999998,"percentFarthestFromBall":40.58687499999999},"demo":{"inflicted":1,"taken":1}},"advanced":{"goalParticipation":60,"rating":0.584025439789442}}]},"number":9,"games":[{"_id":"6082fb4c0d9dcf9da5a4d2ea","blue":1,"orange":5,"duration":300,"ballchasing":"dd308bf5-5678-4e79-96ea-c937dc631b41"},{"_id":"6082fb4d0d9dcf9da5a4d2f1","blue":1,"orange":0,"duration":300,"ballchasing":"145598b6-3116-40f5-a74f-7fa0b5b3dabf"},{"_id":"6082fb4d0d9dcf9da5a4d2f8","blue":3,"orange":0,"duration":300,"ballchasing":"d409cce2-e03d-4f5b-9e8f-63507e976adb"},{"_id":"6082fb4e0d9dcf9da5a4d2ff","blue":1,"orange":0,"duration":300,"ballchasing":"4bec9148-fbac-4b93-b500-6730f7ae833a"}]}
  
## Match [/matches/{id}/games]

- Parameters
  + id: 6043152fa09e7fba40d2ae62 (required, string) - a match id

### Get Match Games [GET]

- Response 200 (application/json)
{"games":[{"_id":"6082fb4c0d9dcf9da5a4d2ea","octane_id":"0350109","number":1,"match":{"_id":"6043152fa09e7fba40d2ae62","slug":"ae62-flipsid3-tactics-vs-nrg-esports","event":{"_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","region":"INT","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"]},"stage":{"_id":0,"name":"Main Event","format":"bracket-8de"},"format":{"type":"best","length":5}},"map":{"id":"stadium_p","name":"DFH Stadium"},"duration":300,"date":"2016-12-04T13:28:00Z","blue":{"matchWinner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"},"stats":{"ball":{"possessionTime":151.1,"timeInSide":132},"core":{"shots":6,"goals":1,"saves":4,"assists":1,"score":555,"shootingPercentage":16.666666666666664},"boost":{"bpm":1206,"bcpm":1254.6616199999999,"avgAmount":149.01,"amountCollected":7166,"amountStolen":1121,"amountCollectedBig":5224,"amountStolenBig":554,"amountCollectedSmall":1942,"amountStolenSmall":567,"countCollectedBig":61,"countStolenBig":8,"countCollectedSmall":174,"countStolenSmall":46,"amountOverfill":845,"amountOverfillStolen":155,"amountUsedWhileSupersonic":1099,"timeZeroBoost":95.04,"timeFullBoost":140.88,"timeBoost0To25":351.61,"timeBoost25To50":243.24,"timeBoost50To75":197.32999999999998,"timeBoost75To100":255.28},"movement":{"totalDistance":1564432,"timeSupersonicSpeed":155.48000000000002,"timeBoostSpeed":454.13,"timeSlowSpeed":468.14,"timeGround":647.07,"timeLowAir":388.42,"timeHighAir":42.26,"timePowerslide":23.35,"countPowerslide":198},"positioning":{"timeDefensiveThird":489.67999999999995,"timeNeutralThird":372.57,"timeOffensiveThird":215.48999999999998,"timeDefensiveHalf":687.13,"timeOffensiveHalf":390.42,"timeBehindBall":790.68,"timeInfrontBall":287.06},"demo":{"inflicted":0,"taken":3}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d98","slug":"3d98-markydooda","tag":"Markydooda","country":"ab"},"stats":{"core":{"shots":3,"goals":0,"saves":1,"assists":1,"score":190,"shootingPercentage":0},"boost":{"bpm":416,"bcpm":432.46082,"avgAmount":49.79,"amountCollected":2470,"amountStolen":309,"amountCollectedBig":1938,"amountStolenBig":150,"amountCollectedSmall":532,"amountStolenSmall":159,"countCollectedBig":23,"countStolenBig":2,"countCollectedSmall":49,"countStolenSmall":15,"amountOverfill":367,"amountOverfillStolen":49,"amountUsedWhileSupersonic":463,"timeZeroBoost":41.59,"percentZeroBoost":12.136333,"timeFullBoost":60.49,"percentFullBoost":17.651522,"timeBoost0To25":108.91,"timeBoost25To50":77.91,"timeBoost50To75":62.83,"timeBoost75To100":92.27,"percentBoost0To25":31.852478,"percentBoost25To50":22.786032,"percentBoost50To75":18.375643,"percentBoost75To100":26.985844},"movement":{"avgSpeed":1564,"totalDistance":522243,"timeSupersonicSpeed":53.11,"timeBoostSpeed":147.86,"timeSlowSpeed":158.81,"timeGround":206.81,"timeLowAir":136.39,"timeHighAir":16.58,"timePowerslide":7.49,"countPowerslide":78,"avgPowerslideDuration":0.1,"avgSpeedPercentage":68,"percentSlowSpeed":44.14087,"percentBoostSpeed":41.09734,"percentSupersonicSpeed":14.7618,"percentGround":57.48235,"percentLowAir":37.90928,"percentHighAir":4.6083717},"positioning":{"avgDistanceToBall":3237,"avgDistanceToBallPossession":3118,"avgDistanceToBallNoPossession":3380,"avgDistanceToMates":4099,"timeDefensiveThird":167.34,"timeNeutralThird":107.81,"timeOffensiveThird":84.63,"timeDefensiveHalf":220.75,"timeOffensiveHalf":138.98,"timeBehindBall":255.76,"timeInfrontBall":104.02,"timeMostBack":111.3,"timeMostForward":112,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":132.8,"timeFarthestFromBall":115.3,"percentDefensiveThird":46.511757,"percentOffensiveThird":23.522709,"percentNeutralThird":29.965534,"percentDefensiveHalf":61.36547,"percentOffensiveHalf":38.634533,"percentBehindBall":71.08789,"percentInfrontBall":28.912113,"percentMostBack":32.478333,"percentMostForward":32.6826,"percentClosestToBall":38.752224,"percentFarthestFromBall":33.64557},"demo":{"inflicted":0,"taken":1}},"advanced":{"goalParticipation":100,"rating":0.904279163121216}},{"player":{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","country":"it"},"stats":{"core":{"shots":1,"goals":1,"saves":2,"assists":0,"score":185,"shootingPercentage":100},"boost":{"bpm":339,"bcpm":369.95535,"avgAmount":52.79,"amountCollected":2113,"amountStolen":439,"amountCollectedBig":1469,"amountStolenBig":249,"amountCollectedSmall":644,"amountStolenSmall":190,"countCollectedBig":18,"countStolenBig":4,"countCollectedSmall":58,"countStolenSmall":16,"amountOverfill":268,"amountOverfillStolen":62,"amountUsedWhileSupersonic":272,"timeZeroBoost":17.03,"percentZeroBoost":4.9695063,"timeFullBoost":43.41,"percentFullBoost":12.667425,"timeBoost0To25":106.08,"timeBoost25To50":79.96,"timeBoost50To75":83.32,"timeBoost75To100":85.67,"percentBoost0To25":29.879162,"percentBoost25To50":22.52204,"percentBoost50To75":23.468437,"percentBoost75To100":24.130354},"movement":{"avgSpeed":1475,"totalDistance":497882,"timeSupersonicSpeed":33.63,"timeBoostSpeed":152.93,"timeSlowSpeed":176.37,"timeGround":224.41,"timeLowAir":129.36,"timeHighAir":9.16,"timePowerslide":7.16,"countPowerslide":55,"avgPowerslideDuration":0.13,"avgSpeedPercentage":64.13043,"percentSlowSpeed":48.59615,"percentBoostSpeed":42.1376,"percentSupersonicSpeed":9.26625,"percentGround":61.832855,"percentLowAir":35.643234,"percentHighAir":2.5239024},"positioning":{"avgDistanceToBall":3370,"avgDistanceToBallPossession":3206,"avgDistanceToBallNoPossession":3548,"avgDistanceToMates":3956,"timeDefensiveThird":165.57,"timeNeutralThird":138.46,"timeOffensiveThird":58.89,"timeDefensiveHalf":240,"timeOffensiveHalf":122.83,"timeBehindBall":274.64,"timeInfrontBall":88.28,"timeMostBack":135.8,"timeMostForward":105.5,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":108.4,"timeFarthestFromBall":117.5,"percentDefensiveThird":45.621624,"percentOffensiveThird":16.226715,"percentNeutralThird":38.15166,"percentDefensiveHalf":66.146675,"percentOffensiveHalf":33.853317,"percentBehindBall":75.67508,"percentInfrontBall":24.324919,"percentMostBack":39.62765,"percentMostForward":30.785841,"percentClosestToBall":31.632088,"percentFarthestFromBall":34.287548},"demo":{"inflicted":0,"taken":0}},"advanced":{"goalParticipation":100,"rating":1.2740661980234818}},{"player":{"_id":"6082fb1d0d9dcf9da5a4d079","slug":"d079-greazymeister","tag":"gReazymeister"},"stats":{"core":{"shots":2,"goals":0,"saves":1,"assists":0,"score":180,"shootingPercentage":0},"boost":{"bpm":451,"bcpm":452.24545,"avgAmount":46.43,"amountCollected":2583,"amountStolen":373,"amountCollectedBig":1817,"amountStolenBig":155,"amountCollectedSmall":766,"amountStolenSmall":218,"countCollectedBig":20,"countStolenBig":2,"countCollectedSmall":67,"countStolenSmall":15,"amountOverfill":210,"amountOverfillStolen":44,"amountUsedWhileSupersonic":364,"timeZeroBoost":36.42,"percentZeroBoost":10.62768,"timeFullBoost":36.98,"percentFullBoost":10.791094,"timeBoost0To25":136.62,"timeBoost25To50":85.37,"timeBoost50To75":51.18,"timeBoost75To100":77.34,"percentBoost0To25":38.977493,"percentBoost25To50":24.35594,"percentBoost50To75":14.601582,"percentBoost75To100":22.064991},"movement":{"avgSpeed":1651,"totalDistance":544307,"timeSupersonicSpeed":68.74,"timeBoostSpeed":153.34,"timeSlowSpeed":132.96,"timeGround":215.85,"timeLowAir":122.67,"timeHighAir":16.52,"timePowerslide":8.7,"countPowerslide":65,"avgPowerslideDuration":0.13,"avgSpeedPercentage":71.78261,"percentSlowSpeed":37.449306,"percentBoostSpeed":43.189503,"percentSupersonicSpeed":19.3612,"percentGround":60.795963,"percentLowAir":34.551037,"percentHighAir":4.6529965},"positioning":{"avgDistanceToBall":3449,"avgDistanceToBallPossession":3244,"avgDistanceToBallNoPossession":3657,"avgDistanceToMates":4104,"timeDefensiveThird":156.77,"timeNeutralThird":126.3,"timeOffensiveThird":71.97,"timeDefensiveHalf":226.38,"timeOffensiveHalf":128.61,"timeBehindBall":260.28,"timeInfrontBall":94.76,"timeMostBack":115.4,"timeMostForward":133.2,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":109.8,"timeFarthestFromBall":129.7,"percentDefensiveThird":44.155586,"percentOffensiveThird":20.270954,"percentNeutralThird":35.573456,"percentDefensiveHalf":63.770813,"percentOffensiveHalf":36.22919,"percentBehindBall":73.31005,"percentInfrontBall":26.689949,"percentMostBack":33.67475,"percentMostForward":38.86895,"percentClosestToBall":32.04062,"percentFarthestFromBall":37.84762},"demo":{"inflicted":0,"taken":2}},"advanced":{"goalParticipation":0,"rating":0.3056871660876292}}]},"orange":{"winner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023a0","slug":"23a0-nrg-esports","name":"NRG Esports","image":"https://griffon.octane.gg/teams/nrg-esports.png"},"stats":{"ball":{"possessionTime":129.33,"timeInSide":200.98},"core":{"shots":10,"goals":5,"saves":5,"assists":3,"score":995,"shootingPercentage":50},"boost":{"bpm":1262,"bcpm":1328.3726,"avgAmount":143.97,"amountCollected":7587,"amountStolen":1427,"amountCollectedBig":5641,"amountStolenBig":1028,"amountCollectedSmall":1946,"amountStolenSmall":399,"countCollectedBig":65,"countStolenBig":12,"countCollectedSmall":180,"countStolenSmall":38,"amountOverfill":885,"amountOverfillStolen":184,"amountUsedWhileSupersonic":1471,"timeZeroBoost":126.86,"timeFullBoost":133,"timeBoost0To25":368.98,"timeBoost25To50":227.51,"timeBoost50To75":199.1,"timeBoost75To100":247.36},"movement":{"totalDistance":1545937,"timeSupersonicSpeed":158.03,"timeBoostSpeed":435.11,"timeSlowSpeed":495.15999999999997,"timeGround":653.49,"timeLowAir":387.39,"timeHighAir":47.410000000000004,"timePowerslide":20.439999999999998,"countPowerslide":139},"positioning":{"timeDefensiveThird":614.79,"timeNeutralThird":313.28,"timeOffensiveThird":160.20999999999998,"timeDefensiveHalf":790.5699999999999,"timeOffensiveHalf":297.67,"timeBehindBall":818.3599999999999,"timeInfrontBall":269.93},"demo":{"inflicted":3,"taken":0}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7b","slug":"3d7b-jacob","tag":"Jacob","country":"us"},"stats":{"core":{"shots":4,"goals":3,"saves":1,"assists":1,"score":405,"shootingPercentage":75},"boost":{"bpm":454,"bcpm":453.296,"avgAmount":48.64,"amountCollected":2589,"amountStolen":562,"amountCollectedBig":1857,"amountStolenBig":438,"amountCollectedSmall":732,"amountStolenSmall":124,"countCollectedBig":21,"countStolenBig":5,"countCollectedSmall":66,"countStolenSmall":12,"amountOverfill":244,"amountOverfillStolen":61,"amountUsedWhileSupersonic":470,"timeZeroBoost":28.89,"percentZeroBoost":8.43036,"timeFullBoost":39.65,"percentFullBoost":11.570225,"timeBoost0To25":101.89,"timeBoost25To50":102.34,"timeBoost50To75":68.03,"timeBoost75To100":73.69,"percentBoost0To25":29.452232,"percentBoost25To50":29.582308,"percentBoost50To75":19.66469,"percentBoost75To100":21.300766},"movement":{"avgSpeed":1607,"totalDistance":543261,"timeSupersonicSpeed":51.22,"timeBoostSpeed":169.48,"timeSlowSpeed":142.22,"timeGround":218.63,"timeLowAir":125.19,"timeHighAir":19.1,"timePowerslide":9.37,"countPowerslide":58,"avgPowerslideDuration":0.16,"avgSpeedPercentage":69.86957,"percentSlowSpeed":39.1877,"percentBoostSpeed":46.698994,"percentSupersonicSpeed":14.113302,"percentGround":60.241924,"percentLowAir":34.495205,"percentHighAir":5.2628675},"positioning":{"avgDistanceToBall":2988,"avgDistanceToBallPossession":2684,"avgDistanceToBallNoPossession":3288,"avgDistanceToMates":3611,"timeDefensiveThird":182.59,"timeNeutralThird":113.26,"timeOffensiveThird":67.07,"timeDefensiveHalf":243.16,"timeOffensiveHalf":119.71,"timeBehindBall":254.35,"timeInfrontBall":108.57,"timeMostBack":77.5,"timeMostForward":143.6,"goalsAgainstWhileLastDefender":0,"timeClosestToBall":115.5,"timeFarthestFromBall":110.9,"percentDefensiveThird":50.311363,"percentOffensiveThird":18.480656,"percentNeutralThird":31.20798,"percentDefensiveHalf":67.01022,"percentOffensiveHalf":32.989777,"percentBehindBall":70.08431,"percentInfrontBall":29.915682,"percentMostBack":22.61519,"percentMostForward":41.903763,"percentClosestToBall":33.70393,"percentFarthestFromBall":32.36161},"demo":{"inflicted":1,"taken":0}},"advanced":{"goalParticipation":80,"rating":2.067937328529509,"mvp":true}},{"player":{"_id":"5f3d8fdd95f40596eae23d7c","slug":"3d7c-sadjunior","tag":"Sadjunior","country":"ca"},"stats":{"core":{"shots":3,"goals":1,"saves":1,"assists":2,"score":335,"shootingPercentage":33.33333333333333},"boost":{"bpm":375,"bcpm":414.42703,"avgAmount":43.46,"amountCollected":2367,"amountStolen":198,"amountCollectedBig":1772,"amountStolenBig":88,"amountCollectedSmall":595,"amountStolenSmall":110,"countCollectedBig":20,"countStolenBig":1,"countCollectedSmall":57,"countStolenSmall":10,"amountOverfill":225,"amountOverfillStolen":11,"amountUsedWhileSupersonic":293,"timeZeroBoost":68.28,"percentZeroBoost":19.924713,"timeFullBoost":32.51,"percentFullBoost":9.486708,"timeBoost0To25":149.62,"timeBoost25To50":62.88,"timeBoost50To75":46.82,"timeBoost75To100":87.03,"percentBoost0To25":43.199074,"percentBoost25To50":18.155045,"percentBoost50To75":13.518117,"percentBoost75To100":25.12776},"movement":{"avgSpeed":1483,"totalDistance":501336,"timeSupersonicSpeed":49.28,"timeBoostSpeed":138.15,"timeSlowSpeed":175.6,"timeGround":204.1,"timeLowAir":135.29,"timeHighAir":23.63,"timePowerslide":4.4,"countPowerslide":33,"avgPowerslideDuration":0.13,"avgSpeedPercentage":64.478264,"percentSlowSpeed":48.37066,"percentBoostSpeed":38.054703,"percentSupersonicSpeed":13.5746355,"percentGround":56.222794,"percentLowAir":37.267914,"percentHighAir":6.509283},"positioning":{"avgDistanceToBall":3203,"avgDistanceToBallPossession":3174,"avgDistanceToBallNoPossession":3239,"avgDistanceToMates":3731,"timeDefensiveThird":217.72,"timeNeutralThird":108.39,"timeOffensiveThird":36.91,"timeDefensiveHalf":279.57,"timeOffensiveHalf":83.45,"timeBehindBall":290.67,"timeInfrontBall":72.35,"timeMostBack":139.4,"timeMostForward":101.3,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":116.5,"timeFarthestFromBall":133.3,"percentDefensiveThird":59.974655,"percentOffensiveThird":10.167483,"percentNeutralThird":29.857857,"percentDefensiveHalf":77.01228,"percentOffensiveHalf":22.987713,"percentBehindBall":80.06997,"percentInfrontBall":19.93003,"percentMostBack":40.67816,"percentMostForward":29.560244,"percentClosestToBall":33.99574,"percentFarthestFromBall":38.89813},"demo":{"inflicted":0,"taken":0}},"advanced":{"goalParticipation":60,"rating":1.5189731338417782}},{"player":{"_id":"5f3d8fdd95f40596eae23d7a","slug":"3d7a-fireburner","tag":"Fireburner","country":"us"},"stats":{"core":{"shots":3,"goals":1,"saves":3,"assists":0,"score":255,"shootingPercentage":33.33333333333333},"boost":{"bpm":433,"bcpm":460.64957,"avgAmount":51.87,"amountCollected":2631,"amountStolen":667,"amountCollectedBig":2012,"amountStolenBig":502,"amountCollectedSmall":619,"amountStolenSmall":165,"countCollectedBig":24,"countStolenBig":6,"countCollectedSmall":57,"countStolenSmall":16,"amountOverfill":416,"amountOverfillStolen":112,"amountUsedWhileSupersonic":708,"timeZeroBoost":29.69,"percentZeroBoost":8.663807,"timeFullBoost":60.84,"percentFullBoost":17.753654,"timeBoost0To25":117.47,"timeBoost25To50":62.29,"timeBoost50To75":84.25,"timeBoost75To100":86.64,"percentBoost0To25":33.50064,"percentBoost25To50":17.76415,"percentBoost50To75":24.026806,"percentBoost75To100":24.708397},"movement":{"avgSpeed":1487,"totalDistance":501340,"timeSupersonicSpeed":57.53,"timeBoostSpeed":127.48,"timeSlowSpeed":177.34,"timeGround":230.76,"timeLowAir":126.91,"timeHighAir":4.68,"timePowerslide":6.67,"countPowerslide":48,"avgPowerslideDuration":0.14,"avgSpeedPercentage":64.652176,"percentSlowSpeed":48.94163,"percentBoostSpeed":35.181454,"percentSupersonicSpeed":15.876914,"percentGround":63.684288,"percentLowAir":35.02415,"percentHighAir":1.2915689},"positioning":{"avgDistanceToBall":3056,"avgDistanceToBallPossession":3039,"avgDistanceToBallNoPossession":3056,"avgDistanceToMates":3782,"timeDefensiveThird":214.48,"timeNeutralThird":91.63,"timeOffensiveThird":56.23,"timeDefensiveHalf":267.84,"timeOffensiveHalf":94.51,"timeBehindBall":273.34,"timeInfrontBall":89.01,"timeMostBack":132.7,"timeMostForward":103.8,"goalsAgainstWhileLastDefender":0,"timeClosestToBall":116.7,"timeFarthestFromBall":104.6,"percentDefensiveThird":59.193024,"percentOffensiveThird":15.518574,"percentNeutralThird":25.288403,"percentDefensiveHalf":73.91748,"percentOffensiveHalf":26.082516,"percentBehindBall":75.43535,"percentInfrontBall":24.564648,"percentMostBack":38.723045,"percentMostForward":30.289766,"percentClosestToBall":34.0541,"percentFarthestFromBall":30.523212},"demo":{"inflicted":2,"taken":0}},"advanced":{"goalParticipation":20,"rating":0.9718027887128077}}]},"ballchasing":"dd308bf5-5678-4e79-96ea-c937dc631b41"}]}

# Group Games

## Games [/games]

### List Games [GET]

+ Parameters
    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id
    
    + sort: name (string, optional) - field to sort by
    
    + order: `asc` (enum[string], optional) - order of sort
        + Members
          + asc
          + desc
    
    + page: 1 (int) - page number
        + Default: 1
    
    + perPage: 20 (int) - results per page
        + Default: 50

* Response 200 (application/json)
{"games":[{"_id":"6082fb4c0d9dcf9da5a4d2ea","octane_id":"0350109","number":1,"match":{"_id":"6043152fa09e7fba40d2ae62","slug":"ae62-flipsid3-tactics-vs-nrg-esports","event":{"_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","region":"INT","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"]},"stage":{"_id":0,"name":"Main Event","format":"bracket-8de"},"format":{"type":"best","length":5}},"map":{"id":"stadium_p","name":"DFH Stadium"},"duration":300,"date":"2016-12-04T13:28:00Z","blue":{"matchWinner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"},"stats":{"ball":{"possessionTime":151.1,"timeInSide":132},"core":{"shots":6,"goals":1,"saves":4,"assists":1,"score":555,"shootingPercentage":16.666666666666664},"boost":{"bpm":1206,"bcpm":1254.6616199999999,"avgAmount":149.01,"amountCollected":7166,"amountStolen":1121,"amountCollectedBig":5224,"amountStolenBig":554,"amountCollectedSmall":1942,"amountStolenSmall":567,"countCollectedBig":61,"countStolenBig":8,"countCollectedSmall":174,"countStolenSmall":46,"amountOverfill":845,"amountOverfillStolen":155,"amountUsedWhileSupersonic":1099,"timeZeroBoost":95.04,"timeFullBoost":140.88,"timeBoost0To25":351.61,"timeBoost25To50":243.24,"timeBoost50To75":197.32999999999998,"timeBoost75To100":255.28},"movement":{"totalDistance":1564432,"timeSupersonicSpeed":155.48000000000002,"timeBoostSpeed":454.13,"timeSlowSpeed":468.14,"timeGround":647.07,"timeLowAir":388.42,"timeHighAir":42.26,"timePowerslide":23.35,"countPowerslide":198},"positioning":{"timeDefensiveThird":489.67999999999995,"timeNeutralThird":372.57,"timeOffensiveThird":215.48999999999998,"timeDefensiveHalf":687.13,"timeOffensiveHalf":390.42,"timeBehindBall":790.68,"timeInfrontBall":287.06},"demo":{"inflicted":0,"taken":3}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d98","slug":"3d98-markydooda","tag":"Markydooda","country":"ab"},"stats":{"core":{"shots":3,"goals":0,"saves":1,"assists":1,"score":190,"shootingPercentage":0},"boost":{"bpm":416,"bcpm":432.46082,"avgAmount":49.79,"amountCollected":2470,"amountStolen":309,"amountCollectedBig":1938,"amountStolenBig":150,"amountCollectedSmall":532,"amountStolenSmall":159,"countCollectedBig":23,"countStolenBig":2,"countCollectedSmall":49,"countStolenSmall":15,"amountOverfill":367,"amountOverfillStolen":49,"amountUsedWhileSupersonic":463,"timeZeroBoost":41.59,"percentZeroBoost":12.136333,"timeFullBoost":60.49,"percentFullBoost":17.651522,"timeBoost0To25":108.91,"timeBoost25To50":77.91,"timeBoost50To75":62.83,"timeBoost75To100":92.27,"percentBoost0To25":31.852478,"percentBoost25To50":22.786032,"percentBoost50To75":18.375643,"percentBoost75To100":26.985844},"movement":{"avgSpeed":1564,"totalDistance":522243,"timeSupersonicSpeed":53.11,"timeBoostSpeed":147.86,"timeSlowSpeed":158.81,"timeGround":206.81,"timeLowAir":136.39,"timeHighAir":16.58,"timePowerslide":7.49,"countPowerslide":78,"avgPowerslideDuration":0.1,"avgSpeedPercentage":68,"percentSlowSpeed":44.14087,"percentBoostSpeed":41.09734,"percentSupersonicSpeed":14.7618,"percentGround":57.48235,"percentLowAir":37.90928,"percentHighAir":4.6083717},"positioning":{"avgDistanceToBall":3237,"avgDistanceToBallPossession":3118,"avgDistanceToBallNoPossession":3380,"avgDistanceToMates":4099,"timeDefensiveThird":167.34,"timeNeutralThird":107.81,"timeOffensiveThird":84.63,"timeDefensiveHalf":220.75,"timeOffensiveHalf":138.98,"timeBehindBall":255.76,"timeInfrontBall":104.02,"timeMostBack":111.3,"timeMostForward":112,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":132.8,"timeFarthestFromBall":115.3,"percentDefensiveThird":46.511757,"percentOffensiveThird":23.522709,"percentNeutralThird":29.965534,"percentDefensiveHalf":61.36547,"percentOffensiveHalf":38.634533,"percentBehindBall":71.08789,"percentInfrontBall":28.912113,"percentMostBack":32.478333,"percentMostForward":32.6826,"percentClosestToBall":38.752224,"percentFarthestFromBall":33.64557},"demo":{"inflicted":0,"taken":1}},"advanced":{"goalParticipation":100,"rating":0.904279163121216}},{"player":{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","country":"it"},"stats":{"core":{"shots":1,"goals":1,"saves":2,"assists":0,"score":185,"shootingPercentage":100},"boost":{"bpm":339,"bcpm":369.95535,"avgAmount":52.79,"amountCollected":2113,"amountStolen":439,"amountCollectedBig":1469,"amountStolenBig":249,"amountCollectedSmall":644,"amountStolenSmall":190,"countCollectedBig":18,"countStolenBig":4,"countCollectedSmall":58,"countStolenSmall":16,"amountOverfill":268,"amountOverfillStolen":62,"amountUsedWhileSupersonic":272,"timeZeroBoost":17.03,"percentZeroBoost":4.9695063,"timeFullBoost":43.41,"percentFullBoost":12.667425,"timeBoost0To25":106.08,"timeBoost25To50":79.96,"timeBoost50To75":83.32,"timeBoost75To100":85.67,"percentBoost0To25":29.879162,"percentBoost25To50":22.52204,"percentBoost50To75":23.468437,"percentBoost75To100":24.130354},"movement":{"avgSpeed":1475,"totalDistance":497882,"timeSupersonicSpeed":33.63,"timeBoostSpeed":152.93,"timeSlowSpeed":176.37,"timeGround":224.41,"timeLowAir":129.36,"timeHighAir":9.16,"timePowerslide":7.16,"countPowerslide":55,"avgPowerslideDuration":0.13,"avgSpeedPercentage":64.13043,"percentSlowSpeed":48.59615,"percentBoostSpeed":42.1376,"percentSupersonicSpeed":9.26625,"percentGround":61.832855,"percentLowAir":35.643234,"percentHighAir":2.5239024},"positioning":{"avgDistanceToBall":3370,"avgDistanceToBallPossession":3206,"avgDistanceToBallNoPossession":3548,"avgDistanceToMates":3956,"timeDefensiveThird":165.57,"timeNeutralThird":138.46,"timeOffensiveThird":58.89,"timeDefensiveHalf":240,"timeOffensiveHalf":122.83,"timeBehindBall":274.64,"timeInfrontBall":88.28,"timeMostBack":135.8,"timeMostForward":105.5,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":108.4,"timeFarthestFromBall":117.5,"percentDefensiveThird":45.621624,"percentOffensiveThird":16.226715,"percentNeutralThird":38.15166,"percentDefensiveHalf":66.146675,"percentOffensiveHalf":33.853317,"percentBehindBall":75.67508,"percentInfrontBall":24.324919,"percentMostBack":39.62765,"percentMostForward":30.785841,"percentClosestToBall":31.632088,"percentFarthestFromBall":34.287548},"demo":{"inflicted":0,"taken":0}},"advanced":{"goalParticipation":100,"rating":1.2740661980234818}},{"player":{"_id":"6082fb1d0d9dcf9da5a4d079","slug":"d079-greazymeister","tag":"gReazymeister"},"stats":{"core":{"shots":2,"goals":0,"saves":1,"assists":0,"score":180,"shootingPercentage":0},"boost":{"bpm":451,"bcpm":452.24545,"avgAmount":46.43,"amountCollected":2583,"amountStolen":373,"amountCollectedBig":1817,"amountStolenBig":155,"amountCollectedSmall":766,"amountStolenSmall":218,"countCollectedBig":20,"countStolenBig":2,"countCollectedSmall":67,"countStolenSmall":15,"amountOverfill":210,"amountOverfillStolen":44,"amountUsedWhileSupersonic":364,"timeZeroBoost":36.42,"percentZeroBoost":10.62768,"timeFullBoost":36.98,"percentFullBoost":10.791094,"timeBoost0To25":136.62,"timeBoost25To50":85.37,"timeBoost50To75":51.18,"timeBoost75To100":77.34,"percentBoost0To25":38.977493,"percentBoost25To50":24.35594,"percentBoost50To75":14.601582,"percentBoost75To100":22.064991},"movement":{"avgSpeed":1651,"totalDistance":544307,"timeSupersonicSpeed":68.74,"timeBoostSpeed":153.34,"timeSlowSpeed":132.96,"timeGround":215.85,"timeLowAir":122.67,"timeHighAir":16.52,"timePowerslide":8.7,"countPowerslide":65,"avgPowerslideDuration":0.13,"avgSpeedPercentage":71.78261,"percentSlowSpeed":37.449306,"percentBoostSpeed":43.189503,"percentSupersonicSpeed":19.3612,"percentGround":60.795963,"percentLowAir":34.551037,"percentHighAir":4.6529965},"positioning":{"avgDistanceToBall":3449,"avgDistanceToBallPossession":3244,"avgDistanceToBallNoPossession":3657,"avgDistanceToMates":4104,"timeDefensiveThird":156.77,"timeNeutralThird":126.3,"timeOffensiveThird":71.97,"timeDefensiveHalf":226.38,"timeOffensiveHalf":128.61,"timeBehindBall":260.28,"timeInfrontBall":94.76,"timeMostBack":115.4,"timeMostForward":133.2,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":109.8,"timeFarthestFromBall":129.7,"percentDefensiveThird":44.155586,"percentOffensiveThird":20.270954,"percentNeutralThird":35.573456,"percentDefensiveHalf":63.770813,"percentOffensiveHalf":36.22919,"percentBehindBall":73.31005,"percentInfrontBall":26.689949,"percentMostBack":33.67475,"percentMostForward":38.86895,"percentClosestToBall":32.04062,"percentFarthestFromBall":37.84762},"demo":{"inflicted":0,"taken":2}},"advanced":{"goalParticipation":0,"rating":0.3056871660876292}}]},"orange":{"winner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023a0","slug":"23a0-nrg-esports","name":"NRG Esports","image":"https://griffon.octane.gg/teams/nrg-esports.png"},"stats":{"ball":{"possessionTime":129.33,"timeInSide":200.98},"core":{"shots":10,"goals":5,"saves":5,"assists":3,"score":995,"shootingPercentage":50},"boost":{"bpm":1262,"bcpm":1328.3726,"avgAmount":143.97,"amountCollected":7587,"amountStolen":1427,"amountCollectedBig":5641,"amountStolenBig":1028,"amountCollectedSmall":1946,"amountStolenSmall":399,"countCollectedBig":65,"countStolenBig":12,"countCollectedSmall":180,"countStolenSmall":38,"amountOverfill":885,"amountOverfillStolen":184,"amountUsedWhileSupersonic":1471,"timeZeroBoost":126.86,"timeFullBoost":133,"timeBoost0To25":368.98,"timeBoost25To50":227.51,"timeBoost50To75":199.1,"timeBoost75To100":247.36},"movement":{"totalDistance":1545937,"timeSupersonicSpeed":158.03,"timeBoostSpeed":435.11,"timeSlowSpeed":495.15999999999997,"timeGround":653.49,"timeLowAir":387.39,"timeHighAir":47.410000000000004,"timePowerslide":20.439999999999998,"countPowerslide":139},"positioning":{"timeDefensiveThird":614.79,"timeNeutralThird":313.28,"timeOffensiveThird":160.20999999999998,"timeDefensiveHalf":790.5699999999999,"timeOffensiveHalf":297.67,"timeBehindBall":818.3599999999999,"timeInfrontBall":269.93},"demo":{"inflicted":3,"taken":0}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7b","slug":"3d7b-jacob","tag":"Jacob","country":"us"},"stats":{"core":{"shots":4,"goals":3,"saves":1,"assists":1,"score":405,"shootingPercentage":75},"boost":{"bpm":454,"bcpm":453.296,"avgAmount":48.64,"amountCollected":2589,"amountStolen":562,"amountCollectedBig":1857,"amountStolenBig":438,"amountCollectedSmall":732,"amountStolenSmall":124,"countCollectedBig":21,"countStolenBig":5,"countCollectedSmall":66,"countStolenSmall":12,"amountOverfill":244,"amountOverfillStolen":61,"amountUsedWhileSupersonic":470,"timeZeroBoost":28.89,"percentZeroBoost":8.43036,"timeFullBoost":39.65,"percentFullBoost":11.570225,"timeBoost0To25":101.89,"timeBoost25To50":102.34,"timeBoost50To75":68.03,"timeBoost75To100":73.69,"percentBoost0To25":29.452232,"percentBoost25To50":29.582308,"percentBoost50To75":19.66469,"percentBoost75To100":21.300766},"movement":{"avgSpeed":1607,"totalDistance":543261,"timeSupersonicSpeed":51.22,"timeBoostSpeed":169.48,"timeSlowSpeed":142.22,"timeGround":218.63,"timeLowAir":125.19,"timeHighAir":19.1,"timePowerslide":9.37,"countPowerslide":58,"avgPowerslideDuration":0.16,"avgSpeedPercentage":69.86957,"percentSlowSpeed":39.1877,"percentBoostSpeed":46.698994,"percentSupersonicSpeed":14.113302,"percentGround":60.241924,"percentLowAir":34.495205,"percentHighAir":5.2628675},"positioning":{"avgDistanceToBall":2988,"avgDistanceToBallPossession":2684,"avgDistanceToBallNoPossession":3288,"avgDistanceToMates":3611,"timeDefensiveThird":182.59,"timeNeutralThird":113.26,"timeOffensiveThird":67.07,"timeDefensiveHalf":243.16,"timeOffensiveHalf":119.71,"timeBehindBall":254.35,"timeInfrontBall":108.57,"timeMostBack":77.5,"timeMostForward":143.6,"goalsAgainstWhileLastDefender":0,"timeClosestToBall":115.5,"timeFarthestFromBall":110.9,"percentDefensiveThird":50.311363,"percentOffensiveThird":18.480656,"percentNeutralThird":31.20798,"percentDefensiveHalf":67.01022,"percentOffensiveHalf":32.989777,"percentBehindBall":70.08431,"percentInfrontBall":29.915682,"percentMostBack":22.61519,"percentMostForward":41.903763,"percentClosestToBall":33.70393,"percentFarthestFromBall":32.36161},"demo":{"inflicted":1,"taken":0}},"advanced":{"goalParticipation":80,"rating":2.067937328529509,"mvp":true}},{"player":{"_id":"5f3d8fdd95f40596eae23d7c","slug":"3d7c-sadjunior","tag":"Sadjunior","country":"ca"},"stats":{"core":{"shots":3,"goals":1,"saves":1,"assists":2,"score":335,"shootingPercentage":33.33333333333333},"boost":{"bpm":375,"bcpm":414.42703,"avgAmount":43.46,"amountCollected":2367,"amountStolen":198,"amountCollectedBig":1772,"amountStolenBig":88,"amountCollectedSmall":595,"amountStolenSmall":110,"countCollectedBig":20,"countStolenBig":1,"countCollectedSmall":57,"countStolenSmall":10,"amountOverfill":225,"amountOverfillStolen":11,"amountUsedWhileSupersonic":293,"timeZeroBoost":68.28,"percentZeroBoost":19.924713,"timeFullBoost":32.51,"percentFullBoost":9.486708,"timeBoost0To25":149.62,"timeBoost25To50":62.88,"timeBoost50To75":46.82,"timeBoost75To100":87.03,"percentBoost0To25":43.199074,"percentBoost25To50":18.155045,"percentBoost50To75":13.518117,"percentBoost75To100":25.12776},"movement":{"avgSpeed":1483,"totalDistance":501336,"timeSupersonicSpeed":49.28,"timeBoostSpeed":138.15,"timeSlowSpeed":175.6,"timeGround":204.1,"timeLowAir":135.29,"timeHighAir":23.63,"timePowerslide":4.4,"countPowerslide":33,"avgPowerslideDuration":0.13,"avgSpeedPercentage":64.478264,"percentSlowSpeed":48.37066,"percentBoostSpeed":38.054703,"percentSupersonicSpeed":13.5746355,"percentGround":56.222794,"percentLowAir":37.267914,"percentHighAir":6.509283},"positioning":{"avgDistanceToBall":3203,"avgDistanceToBallPossession":3174,"avgDistanceToBallNoPossession":3239,"avgDistanceToMates":3731,"timeDefensiveThird":217.72,"timeNeutralThird":108.39,"timeOffensiveThird":36.91,"timeDefensiveHalf":279.57,"timeOffensiveHalf":83.45,"timeBehindBall":290.67,"timeInfrontBall":72.35,"timeMostBack":139.4,"timeMostForward":101.3,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":116.5,"timeFarthestFromBall":133.3,"percentDefensiveThird":59.974655,"percentOffensiveThird":10.167483,"percentNeutralThird":29.857857,"percentDefensiveHalf":77.01228,"percentOffensiveHalf":22.987713,"percentBehindBall":80.06997,"percentInfrontBall":19.93003,"percentMostBack":40.67816,"percentMostForward":29.560244,"percentClosestToBall":33.99574,"percentFarthestFromBall":38.89813},"demo":{"inflicted":0,"taken":0}},"advanced":{"goalParticipation":60,"rating":1.5189731338417782}},{"player":{"_id":"5f3d8fdd95f40596eae23d7a","slug":"3d7a-fireburner","tag":"Fireburner","country":"us"},"stats":{"core":{"shots":3,"goals":1,"saves":3,"assists":0,"score":255,"shootingPercentage":33.33333333333333},"boost":{"bpm":433,"bcpm":460.64957,"avgAmount":51.87,"amountCollected":2631,"amountStolen":667,"amountCollectedBig":2012,"amountStolenBig":502,"amountCollectedSmall":619,"amountStolenSmall":165,"countCollectedBig":24,"countStolenBig":6,"countCollectedSmall":57,"countStolenSmall":16,"amountOverfill":416,"amountOverfillStolen":112,"amountUsedWhileSupersonic":708,"timeZeroBoost":29.69,"percentZeroBoost":8.663807,"timeFullBoost":60.84,"percentFullBoost":17.753654,"timeBoost0To25":117.47,"timeBoost25To50":62.29,"timeBoost50To75":84.25,"timeBoost75To100":86.64,"percentBoost0To25":33.50064,"percentBoost25To50":17.76415,"percentBoost50To75":24.026806,"percentBoost75To100":24.708397},"movement":{"avgSpeed":1487,"totalDistance":501340,"timeSupersonicSpeed":57.53,"timeBoostSpeed":127.48,"timeSlowSpeed":177.34,"timeGround":230.76,"timeLowAir":126.91,"timeHighAir":4.68,"timePowerslide":6.67,"countPowerslide":48,"avgPowerslideDuration":0.14,"avgSpeedPercentage":64.652176,"percentSlowSpeed":48.94163,"percentBoostSpeed":35.181454,"percentSupersonicSpeed":15.876914,"percentGround":63.684288,"percentLowAir":35.02415,"percentHighAir":1.2915689},"positioning":{"avgDistanceToBall":3056,"avgDistanceToBallPossession":3039,"avgDistanceToBallNoPossession":3056,"avgDistanceToMates":3782,"timeDefensiveThird":214.48,"timeNeutralThird":91.63,"timeOffensiveThird":56.23,"timeDefensiveHalf":267.84,"timeOffensiveHalf":94.51,"timeBehindBall":273.34,"timeInfrontBall":89.01,"timeMostBack":132.7,"timeMostForward":103.8,"goalsAgainstWhileLastDefender":0,"timeClosestToBall":116.7,"timeFarthestFromBall":104.6,"percentDefensiveThird":59.193024,"percentOffensiveThird":15.518574,"percentNeutralThird":25.288403,"percentDefensiveHalf":73.91748,"percentOffensiveHalf":26.082516,"percentBehindBall":75.43535,"percentInfrontBall":24.564648,"percentMostBack":38.723045,"percentMostForward":30.289766,"percentClosestToBall":34.0541,"percentFarthestFromBall":30.523212},"demo":{"inflicted":2,"taken":0}},"advanced":{"goalParticipation":20,"rating":0.9718027887128077}}]},"ballchasing":"dd308bf5-5678-4e79-96ea-c937dc631b41"}]}

## Game [/games/{id}]

- Parameters
  + id: 6082fb4c0d9dcf9da5a4d2ea (required, string) - a game id

### Get Game [GET]

- Response 200 (application/json)
{"_id":"6082fb4c0d9dcf9da5a4d2ea","octane_id":"0350109","number":1,"match":{"_id":"6043152fa09e7fba40d2ae62","slug":"ae62-flipsid3-tactics-vs-nrg-esports","event":{"_id":"5f35882d53fbbb5894b43040","slug":"3040-rlcs-season-2-world-championship","name":"RLCS Season 2 World Championship","region":"INT","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs19","rlcs19worlds","rlcsworlds","rlcs2"]},"stage":{"_id":0,"name":"Main Event","format":"bracket-8de"},"format":{"type":"best","length":5}},"map":{"id":"stadium_p","name":"DFH Stadium"},"duration":300,"date":"2016-12-04T13:28:00Z","blue":{"matchWinner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"},"stats":{"ball":{"possessionTime":151.1,"timeInSide":132},"core":{"shots":6,"goals":1,"saves":4,"assists":1,"score":555,"shootingPercentage":16.666666666666664},"boost":{"bpm":1206,"bcpm":1254.6616199999999,"avgAmount":149.01,"amountCollected":7166,"amountStolen":1121,"amountCollectedBig":5224,"amountStolenBig":554,"amountCollectedSmall":1942,"amountStolenSmall":567,"countCollectedBig":61,"countStolenBig":8,"countCollectedSmall":174,"countStolenSmall":46,"amountOverfill":845,"amountOverfillStolen":155,"amountUsedWhileSupersonic":1099,"timeZeroBoost":95.04,"timeFullBoost":140.88,"timeBoost0To25":351.61,"timeBoost25To50":243.24,"timeBoost50To75":197.32999999999998,"timeBoost75To100":255.28},"movement":{"totalDistance":1564432,"timeSupersonicSpeed":155.48000000000002,"timeBoostSpeed":454.13,"timeSlowSpeed":468.14,"timeGround":647.07,"timeLowAir":388.42,"timeHighAir":42.26,"timePowerslide":23.35,"countPowerslide":198},"positioning":{"timeDefensiveThird":489.67999999999995,"timeNeutralThird":372.57,"timeOffensiveThird":215.48999999999998,"timeDefensiveHalf":687.13,"timeOffensiveHalf":390.42,"timeBehindBall":790.68,"timeInfrontBall":287.06},"demo":{"inflicted":0,"taken":3}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d98","slug":"3d98-markydooda","tag":"Markydooda","country":"ab"},"stats":{"core":{"shots":3,"goals":0,"saves":1,"assists":1,"score":190,"shootingPercentage":0},"boost":{"bpm":416,"bcpm":432.46082,"avgAmount":49.79,"amountCollected":2470,"amountStolen":309,"amountCollectedBig":1938,"amountStolenBig":150,"amountCollectedSmall":532,"amountStolenSmall":159,"countCollectedBig":23,"countStolenBig":2,"countCollectedSmall":49,"countStolenSmall":15,"amountOverfill":367,"amountOverfillStolen":49,"amountUsedWhileSupersonic":463,"timeZeroBoost":41.59,"percentZeroBoost":12.136333,"timeFullBoost":60.49,"percentFullBoost":17.651522,"timeBoost0To25":108.91,"timeBoost25To50":77.91,"timeBoost50To75":62.83,"timeBoost75To100":92.27,"percentBoost0To25":31.852478,"percentBoost25To50":22.786032,"percentBoost50To75":18.375643,"percentBoost75To100":26.985844},"movement":{"avgSpeed":1564,"totalDistance":522243,"timeSupersonicSpeed":53.11,"timeBoostSpeed":147.86,"timeSlowSpeed":158.81,"timeGround":206.81,"timeLowAir":136.39,"timeHighAir":16.58,"timePowerslide":7.49,"countPowerslide":78,"avgPowerslideDuration":0.1,"avgSpeedPercentage":68,"percentSlowSpeed":44.14087,"percentBoostSpeed":41.09734,"percentSupersonicSpeed":14.7618,"percentGround":57.48235,"percentLowAir":37.90928,"percentHighAir":4.6083717},"positioning":{"avgDistanceToBall":3237,"avgDistanceToBallPossession":3118,"avgDistanceToBallNoPossession":3380,"avgDistanceToMates":4099,"timeDefensiveThird":167.34,"timeNeutralThird":107.81,"timeOffensiveThird":84.63,"timeDefensiveHalf":220.75,"timeOffensiveHalf":138.98,"timeBehindBall":255.76,"timeInfrontBall":104.02,"timeMostBack":111.3,"timeMostForward":112,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":132.8,"timeFarthestFromBall":115.3,"percentDefensiveThird":46.511757,"percentOffensiveThird":23.522709,"percentNeutralThird":29.965534,"percentDefensiveHalf":61.36547,"percentOffensiveHalf":38.634533,"percentBehindBall":71.08789,"percentInfrontBall":28.912113,"percentMostBack":32.478333,"percentMostForward":32.6826,"percentClosestToBall":38.752224,"percentFarthestFromBall":33.64557},"demo":{"inflicted":0,"taken":1}},"advanced":{"goalParticipation":100,"rating":0.904279163121216}},{"player":{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","country":"it"},"stats":{"core":{"shots":1,"goals":1,"saves":2,"assists":0,"score":185,"shootingPercentage":100},"boost":{"bpm":339,"bcpm":369.95535,"avgAmount":52.79,"amountCollected":2113,"amountStolen":439,"amountCollectedBig":1469,"amountStolenBig":249,"amountCollectedSmall":644,"amountStolenSmall":190,"countCollectedBig":18,"countStolenBig":4,"countCollectedSmall":58,"countStolenSmall":16,"amountOverfill":268,"amountOverfillStolen":62,"amountUsedWhileSupersonic":272,"timeZeroBoost":17.03,"percentZeroBoost":4.9695063,"timeFullBoost":43.41,"percentFullBoost":12.667425,"timeBoost0To25":106.08,"timeBoost25To50":79.96,"timeBoost50To75":83.32,"timeBoost75To100":85.67,"percentBoost0To25":29.879162,"percentBoost25To50":22.52204,"percentBoost50To75":23.468437,"percentBoost75To100":24.130354},"movement":{"avgSpeed":1475,"totalDistance":497882,"timeSupersonicSpeed":33.63,"timeBoostSpeed":152.93,"timeSlowSpeed":176.37,"timeGround":224.41,"timeLowAir":129.36,"timeHighAir":9.16,"timePowerslide":7.16,"countPowerslide":55,"avgPowerslideDuration":0.13,"avgSpeedPercentage":64.13043,"percentSlowSpeed":48.59615,"percentBoostSpeed":42.1376,"percentSupersonicSpeed":9.26625,"percentGround":61.832855,"percentLowAir":35.643234,"percentHighAir":2.5239024},"positioning":{"avgDistanceToBall":3370,"avgDistanceToBallPossession":3206,"avgDistanceToBallNoPossession":3548,"avgDistanceToMates":3956,"timeDefensiveThird":165.57,"timeNeutralThird":138.46,"timeOffensiveThird":58.89,"timeDefensiveHalf":240,"timeOffensiveHalf":122.83,"timeBehindBall":274.64,"timeInfrontBall":88.28,"timeMostBack":135.8,"timeMostForward":105.5,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":108.4,"timeFarthestFromBall":117.5,"percentDefensiveThird":45.621624,"percentOffensiveThird":16.226715,"percentNeutralThird":38.15166,"percentDefensiveHalf":66.146675,"percentOffensiveHalf":33.853317,"percentBehindBall":75.67508,"percentInfrontBall":24.324919,"percentMostBack":39.62765,"percentMostForward":30.785841,"percentClosestToBall":31.632088,"percentFarthestFromBall":34.287548},"demo":{"inflicted":0,"taken":0}},"advanced":{"goalParticipation":100,"rating":1.2740661980234818}},{"player":{"_id":"6082fb1d0d9dcf9da5a4d079","slug":"d079-greazymeister","tag":"gReazymeister"},"stats":{"core":{"shots":2,"goals":0,"saves":1,"assists":0,"score":180,"shootingPercentage":0},"boost":{"bpm":451,"bcpm":452.24545,"avgAmount":46.43,"amountCollected":2583,"amountStolen":373,"amountCollectedBig":1817,"amountStolenBig":155,"amountCollectedSmall":766,"amountStolenSmall":218,"countCollectedBig":20,"countStolenBig":2,"countCollectedSmall":67,"countStolenSmall":15,"amountOverfill":210,"amountOverfillStolen":44,"amountUsedWhileSupersonic":364,"timeZeroBoost":36.42,"percentZeroBoost":10.62768,"timeFullBoost":36.98,"percentFullBoost":10.791094,"timeBoost0To25":136.62,"timeBoost25To50":85.37,"timeBoost50To75":51.18,"timeBoost75To100":77.34,"percentBoost0To25":38.977493,"percentBoost25To50":24.35594,"percentBoost50To75":14.601582,"percentBoost75To100":22.064991},"movement":{"avgSpeed":1651,"totalDistance":544307,"timeSupersonicSpeed":68.74,"timeBoostSpeed":153.34,"timeSlowSpeed":132.96,"timeGround":215.85,"timeLowAir":122.67,"timeHighAir":16.52,"timePowerslide":8.7,"countPowerslide":65,"avgPowerslideDuration":0.13,"avgSpeedPercentage":71.78261,"percentSlowSpeed":37.449306,"percentBoostSpeed":43.189503,"percentSupersonicSpeed":19.3612,"percentGround":60.795963,"percentLowAir":34.551037,"percentHighAir":4.6529965},"positioning":{"avgDistanceToBall":3449,"avgDistanceToBallPossession":3244,"avgDistanceToBallNoPossession":3657,"avgDistanceToMates":4104,"timeDefensiveThird":156.77,"timeNeutralThird":126.3,"timeOffensiveThird":71.97,"timeDefensiveHalf":226.38,"timeOffensiveHalf":128.61,"timeBehindBall":260.28,"timeInfrontBall":94.76,"timeMostBack":115.4,"timeMostForward":133.2,"goalsAgainstWhileLastDefender":2,"timeClosestToBall":109.8,"timeFarthestFromBall":129.7,"percentDefensiveThird":44.155586,"percentOffensiveThird":20.270954,"percentNeutralThird":35.573456,"percentDefensiveHalf":63.770813,"percentOffensiveHalf":36.22919,"percentBehindBall":73.31005,"percentInfrontBall":26.689949,"percentMostBack":33.67475,"percentMostForward":38.86895,"percentClosestToBall":32.04062,"percentFarthestFromBall":37.84762},"demo":{"inflicted":0,"taken":2}},"advanced":{"goalParticipation":0,"rating":0.3056871660876292}}]},"orange":{"winner":true,"team":{"team":{"_id":"6020bc70f1e4807cc70023a0","slug":"23a0-nrg-esports","name":"NRG Esports","image":"https://griffon.octane.gg/teams/nrg-esports.png"},"stats":{"ball":{"possessionTime":129.33,"timeInSide":200.98},"core":{"shots":10,"goals":5,"saves":5,"assists":3,"score":995,"shootingPercentage":50},"boost":{"bpm":1262,"bcpm":1328.3726,"avgAmount":143.97,"amountCollected":7587,"amountStolen":1427,"amountCollectedBig":5641,"amountStolenBig":1028,"amountCollectedSmall":1946,"amountStolenSmall":399,"countCollectedBig":65,"countStolenBig":12,"countCollectedSmall":180,"countStolenSmall":38,"amountOverfill":885,"amountOverfillStolen":184,"amountUsedWhileSupersonic":1471,"timeZeroBoost":126.86,"timeFullBoost":133,"timeBoost0To25":368.98,"timeBoost25To50":227.51,"timeBoost50To75":199.1,"timeBoost75To100":247.36},"movement":{"totalDistance":1545937,"timeSupersonicSpeed":158.03,"timeBoostSpeed":435.11,"timeSlowSpeed":495.15999999999997,"timeGround":653.49,"timeLowAir":387.39,"timeHighAir":47.410000000000004,"timePowerslide":20.439999999999998,"countPowerslide":139},"positioning":{"timeDefensiveThird":614.79,"timeNeutralThird":313.28,"timeOffensiveThird":160.20999999999998,"timeDefensiveHalf":790.5699999999999,"timeOffensiveHalf":297.67,"timeBehindBall":818.3599999999999,"timeInfrontBall":269.93},"demo":{"inflicted":3,"taken":0}}},"players":[{"player":{"_id":"5f3d8fdd95f40596eae23d7b","slug":"3d7b-jacob","tag":"Jacob","country":"us"},"stats":{"core":{"shots":4,"goals":3,"saves":1,"assists":1,"score":405,"shootingPercentage":75},"boost":{"bpm":454,"bcpm":453.296,"avgAmount":48.64,"amountCollected":2589,"amountStolen":562,"amountCollectedBig":1857,"amountStolenBig":438,"amountCollectedSmall":732,"amountStolenSmall":124,"countCollectedBig":21,"countStolenBig":5,"countCollectedSmall":66,"countStolenSmall":12,"amountOverfill":244,"amountOverfillStolen":61,"amountUsedWhileSupersonic":470,"timeZeroBoost":28.89,"percentZeroBoost":8.43036,"timeFullBoost":39.65,"percentFullBoost":11.570225,"timeBoost0To25":101.89,"timeBoost25To50":102.34,"timeBoost50To75":68.03,"timeBoost75To100":73.69,"percentBoost0To25":29.452232,"percentBoost25To50":29.582308,"percentBoost50To75":19.66469,"percentBoost75To100":21.300766},"movement":{"avgSpeed":1607,"totalDistance":543261,"timeSupersonicSpeed":51.22,"timeBoostSpeed":169.48,"timeSlowSpeed":142.22,"timeGround":218.63,"timeLowAir":125.19,"timeHighAir":19.1,"timePowerslide":9.37,"countPowerslide":58,"avgPowerslideDuration":0.16,"avgSpeedPercentage":69.86957,"percentSlowSpeed":39.1877,"percentBoostSpeed":46.698994,"percentSupersonicSpeed":14.113302,"percentGround":60.241924,"percentLowAir":34.495205,"percentHighAir":5.2628675},"positioning":{"avgDistanceToBall":2988,"avgDistanceToBallPossession":2684,"avgDistanceToBallNoPossession":3288,"avgDistanceToMates":3611,"timeDefensiveThird":182.59,"timeNeutralThird":113.26,"timeOffensiveThird":67.07,"timeDefensiveHalf":243.16,"timeOffensiveHalf":119.71,"timeBehindBall":254.35,"timeInfrontBall":108.57,"timeMostBack":77.5,"timeMostForward":143.6,"goalsAgainstWhileLastDefender":0,"timeClosestToBall":115.5,"timeFarthestFromBall":110.9,"percentDefensiveThird":50.311363,"percentOffensiveThird":18.480656,"percentNeutralThird":31.20798,"percentDefensiveHalf":67.01022,"percentOffensiveHalf":32.989777,"percentBehindBall":70.08431,"percentInfrontBall":29.915682,"percentMostBack":22.61519,"percentMostForward":41.903763,"percentClosestToBall":33.70393,"percentFarthestFromBall":32.36161},"demo":{"inflicted":1,"taken":0}},"advanced":{"goalParticipation":80,"rating":2.067937328529509,"mvp":true}},{"player":{"_id":"5f3d8fdd95f40596eae23d7c","slug":"3d7c-sadjunior","tag":"Sadjunior","country":"ca"},"stats":{"core":{"shots":3,"goals":1,"saves":1,"assists":2,"score":335,"shootingPercentage":33.33333333333333},"boost":{"bpm":375,"bcpm":414.42703,"avgAmount":43.46,"amountCollected":2367,"amountStolen":198,"amountCollectedBig":1772,"amountStolenBig":88,"amountCollectedSmall":595,"amountStolenSmall":110,"countCollectedBig":20,"countStolenBig":1,"countCollectedSmall":57,"countStolenSmall":10,"amountOverfill":225,"amountOverfillStolen":11,"amountUsedWhileSupersonic":293,"timeZeroBoost":68.28,"percentZeroBoost":19.924713,"timeFullBoost":32.51,"percentFullBoost":9.486708,"timeBoost0To25":149.62,"timeBoost25To50":62.88,"timeBoost50To75":46.82,"timeBoost75To100":87.03,"percentBoost0To25":43.199074,"percentBoost25To50":18.155045,"percentBoost50To75":13.518117,"percentBoost75To100":25.12776},"movement":{"avgSpeed":1483,"totalDistance":501336,"timeSupersonicSpeed":49.28,"timeBoostSpeed":138.15,"timeSlowSpeed":175.6,"timeGround":204.1,"timeLowAir":135.29,"timeHighAir":23.63,"timePowerslide":4.4,"countPowerslide":33,"avgPowerslideDuration":0.13,"avgSpeedPercentage":64.478264,"percentSlowSpeed":48.37066,"percentBoostSpeed":38.054703,"percentSupersonicSpeed":13.5746355,"percentGround":56.222794,"percentLowAir":37.267914,"percentHighAir":6.509283},"positioning":{"avgDistanceToBall":3203,"avgDistanceToBallPossession":3174,"avgDistanceToBallNoPossession":3239,"avgDistanceToMates":3731,"timeDefensiveThird":217.72,"timeNeutralThird":108.39,"timeOffensiveThird":36.91,"timeDefensiveHalf":279.57,"timeOffensiveHalf":83.45,"timeBehindBall":290.67,"timeInfrontBall":72.35,"timeMostBack":139.4,"timeMostForward":101.3,"goalsAgainstWhileLastDefender":1,"timeClosestToBall":116.5,"timeFarthestFromBall":133.3,"percentDefensiveThird":59.974655,"percentOffensiveThird":10.167483,"percentNeutralThird":29.857857,"percentDefensiveHalf":77.01228,"percentOffensiveHalf":22.987713,"percentBehindBall":80.06997,"percentInfrontBall":19.93003,"percentMostBack":40.67816,"percentMostForward":29.560244,"percentClosestToBall":33.99574,"percentFarthestFromBall":38.89813},"demo":{"inflicted":0,"taken":0}},"advanced":{"goalParticipation":60,"rating":1.5189731338417782}},{"player":{"_id":"5f3d8fdd95f40596eae23d7a","slug":"3d7a-fireburner","tag":"Fireburner","country":"us"},"stats":{"core":{"shots":3,"goals":1,"saves":3,"assists":0,"score":255,"shootingPercentage":33.33333333333333},"boost":{"bpm":433,"bcpm":460.64957,"avgAmount":51.87,"amountCollected":2631,"amountStolen":667,"amountCollectedBig":2012,"amountStolenBig":502,"amountCollectedSmall":619,"amountStolenSmall":165,"countCollectedBig":24,"countStolenBig":6,"countCollectedSmall":57,"countStolenSmall":16,"amountOverfill":416,"amountOverfillStolen":112,"amountUsedWhileSupersonic":708,"timeZeroBoost":29.69,"percentZeroBoost":8.663807,"timeFullBoost":60.84,"percentFullBoost":17.753654,"timeBoost0To25":117.47,"timeBoost25To50":62.29,"timeBoost50To75":84.25,"timeBoost75To100":86.64,"percentBoost0To25":33.50064,"percentBoost25To50":17.76415,"percentBoost50To75":24.026806,"percentBoost75To100":24.708397},"movement":{"avgSpeed":1487,"totalDistance":501340,"timeSupersonicSpeed":57.53,"timeBoostSpeed":127.48,"timeSlowSpeed":177.34,"timeGround":230.76,"timeLowAir":126.91,"timeHighAir":4.68,"timePowerslide":6.67,"countPowerslide":48,"avgPowerslideDuration":0.14,"avgSpeedPercentage":64.652176,"percentSlowSpeed":48.94163,"percentBoostSpeed":35.181454,"percentSupersonicSpeed":15.876914,"percentGround":63.684288,"percentLowAir":35.02415,"percentHighAir":1.2915689},"positioning":{"avgDistanceToBall":3056,"avgDistanceToBallPossession":3039,"avgDistanceToBallNoPossession":3056,"avgDistanceToMates":3782,"timeDefensiveThird":214.48,"timeNeutralThird":91.63,"timeOffensiveThird":56.23,"timeDefensiveHalf":267.84,"timeOffensiveHalf":94.51,"timeBehindBall":273.34,"timeInfrontBall":89.01,"timeMostBack":132.7,"timeMostForward":103.8,"goalsAgainstWhileLastDefender":0,"timeClosestToBall":116.7,"timeFarthestFromBall":104.6,"percentDefensiveThird":59.193024,"percentOffensiveThird":15.518574,"percentNeutralThird":25.288403,"percentDefensiveHalf":73.91748,"percentOffensiveHalf":26.082516,"percentBehindBall":75.43535,"percentInfrontBall":24.564648,"percentMostBack":38.723045,"percentMostForward":30.289766,"percentClosestToBall":34.0541,"percentFarthestFromBall":30.523212},"demo":{"inflicted":2,"taken":0}},"advanced":{"goalParticipation":20,"rating":0.9718027887128077}}]},"ballchasing":"dd308bf5-5678-4e79-96ea-c937dc631b41"}

# Group Players

## Players [/players]

### List Players [GET]

+ Parameters
    + tag: Kro (string, optional) - a portion of the player tag

    + country: us (string, optional) - a 2-letter country code

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + sort: name (string, optional) - field to sort by
    
    + order: `asc` (enum[string], optional) - order of sort
        + Members
          + asc
          + desc
    
    + page: 1 (number) - page number
        + Default: 1
    
    + perPage: 20 (number) - results per page
        + Default: 50

* Response 200 (application/json)
  {"players":[{"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","name":"Francesco Cinquemani","country":"it","team":{"_id":"6020c370f1e4807cc702fc9a","slug":"fc9a-wolves-esports","name":"Wolves Esports","region":"EU","image":"https://griffon.octane.gg/teams/wolves-esports.png"},"accounts":[{"platform":"steam","id":"76561198072696308"}]}]}

## Player [/players/{id}]

- Parameters
  + id: 5f35882d53fbbb5894b43040 (required, string) - a player id

### Get Player [GET]

- Response 200 (application/json)
  {"_id":"5f3d8fdd95f40596eae23d97","slug":"3d97-kuxir97","tag":"kuxir97","name":"Francesco Cinquemani","country":"it","team":{"_id":"6020c370f1e4807cc702fc9a","slug":"fc9a-wolves-esports","name":"Wolves Esports","region":"EU","image":"https://griffon.octane.gg/teams/wolves-esports.png"},"accounts":[{"platform":"steam","id":"76561198072696308"}]}

# Group Teams

## Teams [/teams]

### List Teams [GET]

+ Parameters
    + name: Flip (string, optional) - a portion of the team name
    
    + sort: name (string, optional) - field to sort by
    
    + order: `asc` (enum[string], optional) - order of sort
        + Members
          + asc
          + desc
    
    + page: 1 (number) - page number
        + Default: 1
    
    + perPage: 20 (number) - results per page
        + Default: 50

* Response 200 (application/json)
  {"teams":[{"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","region":"EU","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"}]}

## Teams [/teams/active]

- Parameters
  + region: NA (optional, string) - a region

### List Active Teams [GET]

* Response 200 (application/json)
  {"teams":[{"team":{"_id":"6020bc71f1e4807cc7002514","slug":"2514-1ne-esports","name":"1NE eSports","region":"ASIA","image":"https://griffon.octane.gg/teams/1ne-esports.png"},"players":[{"_id":"5f3d8fdd95f40596eae24355","slug":"4355-gon","tag":"Gon","name":"Handy Setiawan","country":"id","team":{"_id":"6020bc71f1e4807cc7002514","slug":"2514-1ne-esports","name":"1NE eSports","region":"ASIA","image":"https://griffon.octane.gg/teams/1ne-esports.png"}},{"_id":"609b903ea07ea8dc37040007","slug":"0007-misty","tag":"Misty","name":"Thrishernn Raaj","country":"my","team":{"_id":"6020bc71f1e4807cc7002514","slug":"2514-1ne-esports","name":"1NE eSports","region":"ASIA","image":"https://griffon.octane.gg/teams/1ne-esports.png"},"accounts":[{"platform":"steam","id":"76561198185792373"}]},{"_id":"6064e22e3036cdb09d774566","slug":"4566-revoir","tag":"Revoir","name":"Dika Utama","country":"id","team":{"_id":"6020bc71f1e4807cc7002514","slug":"2514-1ne-esports","name":"1NE eSports","region":"ASIA","image":"https://griffon.octane.gg/teams/1ne-esports.png"},"accounts":[{"platform":"steam","id":"76561198142293058"}]},{"_id":"5f3d8fdd95f40596eae23f18","slug":"3f18-squirrel","tag":"Squirrel","name":"Jordy Loing","country":"id","team":{"_id":"6020bc71f1e4807cc7002514","slug":"2514-1ne-esports","name":"1NE eSports","region":"ASIA","image":"https://griffon.octane.gg/teams/1ne-esports.png"},"accounts":[{"platform":"steam","id":"76561198223175356"}],"substitute":true}]}]}

## Team [/teams/{id}]

- Parameters
  + id: 6020bc70f1e4807cc70023c7 (required, string) - a team id

### Get Team [GET]

- Response 200 (application/json)
  {"_id":"6020bc70f1e4807cc70023c7","slug":"23c7-flipsid3-tactics","name":"FlipSid3 Tactics","region":"EU","image":"https://griffon.octane.gg/teams/flipsid3-tactics.png"}

# Group Records

## Player Records [/records/players]

### Get Player Records [GET]

+ Parameters
    + type: `game` (enum[string], required) - type of aggregation 
        + Members
          + game
          + series
    
    + stat: `score` (string, required) - stat for records

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"records":[{"game":{"_id":"60431da0a09e7fba40d5fdef","number":2,"match":{"_id":"60431d73a09e7fba40d5ee9a","slug":"ee9a-exotiik-vs-trex","event":{"_id":"5f35882d53fbbb5894b43158","slug":"3158-the-salt-mine-2-europe","name":"The Salt Mine 2 Europe","region":"EU","mode":1,"tier":"B","image":"https://griffon.octane.gg/events/salt-mine.png"},"stage":{"_id":0,"name":"Qualifier 1","format":"bracket","qualifier":true},"format":{"type":"best","length":5}},"map":{"id":"eurostadium_p","name":"Mannfield"},"duration":300,"date":"2020-06-21T18:20:18Z","ballchasing":"1dfe2c4e-8265-4be5-b018-0535bc7979e0"},"team":{"_id":"6020c3bff1e4807cc7031d84","slug":"1d84-exotiik","name":"ExoTiiK"},"opponent":{"_id":"6020c3bff1e4807cc7031d97","slug":"1d97-trex","name":"Trex"},"winner":true,"player":{"_id":"5f3d8fdd95f40596eae24142","slug":"4142-exotiik","tag":"ExoTiiK","country":"fr"},"stat":2489}]}

## Team Records [/records/teams]

### Get Team Records [GET]

+ Parameters
    + type: `game` (enum[string], required) - type of aggregation 
        + Members
          + game
          + series
    
    + stat: `score` (string, required) - stat for records

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"records":[{"team":{"_id":"6020c1d2f1e4807cc7025e1c","slug":"5e1c-rix-gg","name":"Rix.GG","image":"https://griffon.octane.gg/teams/rix-gg.png"},"game":{"_id":"607d7cb6473f172e4c0cf8e7","number":4,"match":{"_id":"6074eda9a2e354f36d800b76","slug":"0b76-rix-gg-vs-bs-competition","event":{"_id":"6074deb130d2d952f6d666d0","slug":"66d0-rlcs-x-spring-europe-regional-3","name":"RLCS X Spring Europe Regional 3","region":"EU","mode":3,"tier":"A","image":"https://griffon.octane.gg/events/rlcs-x.png","groups":["rlcs","rlcsx","rlcsxspring","rlcseu"]},"stage":{"_id":1,"name":"Group Stage","format":"rr-4g5"},"format":{"type":"best","length":5}},"map":{"id":"eurostadium_night_p","name":"Mannfield (Night)"},"duration":1361,"date":"2021-04-16T18:30:43Z","ballchasing":"416caffd-00af-479b-b0e7-888db608b6cf"},"opponent":{"_id":"6074990c8c85ced5f379e9e0","slug":"e9e0-bs-competition","name":"BS+COMPETITION","image":"https://griffon.octane.gg/teams/BS+COMPETITION.png"},"winner":true,"stat":4623}]}

## Game Records [/records/games]

### Get Game Records [GET]

+ Parameters
    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"records":[{"_id":"607d7cb6473f172e4c0cf8e7","number":4,"match":{"_id":"6074eda9a2e354f36d800b76","slug":"0b76-rix-gg-vs-bs-competition","event":{"_id":"6074deb130d2d952f6d666d0","slug":"66d0-rlcs-x-spring-europe-regional-3","name":"RLCS X Spring Europe Regional 3","region":"EU","mode":3,"tier":"A","image":"https://griffon.octane.gg/events/rlcs-x.png","groups":["rlcs","rlcsx","rlcsxspring","rlcseu"]},"stage":{"_id":1,"name":"Group Stage","format":"rr-4g5"},"format":{"type":"best","length":5}},"map":{"id":"eurostadium_night_p","name":"Mannfield (Night)"},"duration":1361,"date":"2021-04-16T18:30:43Z","blue":{"winner":true,"matchWinner":true,"team":{"team":{"_id":"6020c1d2f1e4807cc7025e1c","slug":"5e1c-rix-gg","name":"Rix.GG","image":"https://griffon.octane.gg/teams/rix-gg.png"}}},"orange":{"team":{"team":{"_id":"6074990c8c85ced5f379e9e0","slug":"e9e0-bs-competition","name":"BS+COMPETITION","image":"https://griffon.octane.gg/teams/BS+COMPETITION.png"}}},"stat":8041}]


## Series Records [/records/series]

### Get Series Records [GET]

+ Parameters
    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"records":[{"_id":"607d7cb6473f172e4c0cf8e7","number":4,"match":{"_id":"6074eda9a2e354f36d800b76","slug":"0b76-rix-gg-vs-bs-competition","event":{"_id":"6074deb130d2d952f6d666d0","slug":"66d0-rlcs-x-spring-europe-regional-3","name":"RLCS X Spring Europe Regional 3","region":"EU","mode":3,"tier":"A","image":"https://griffon.octane.gg/events/rlcs-x.png","groups":["rlcs","rlcsx","rlcsxspring","rlcseu"]},"stage":{"_id":1,"name":"Group Stage","format":"rr-4g5"},"format":{"type":"best","length":5}},"map":{"id":"eurostadium_night_p","name":"Mannfield (Night)"},"duration":1361,"date":"2021-04-16T18:30:43Z","blue":{"winner":true,"matchWinner":true,"team":{"team":{"_id":"6020c1d2f1e4807cc7025e1c","slug":"5e1c-rix-gg","name":"Rix.GG","image":"https://griffon.octane.gg/teams/rix-gg.png"}}},"orange":{"team":{"team":{"_id":"6074990c8c85ced5f379e9e0","slug":"e9e0-bs-competition","name":"BS+COMPETITION","image":"https://griffon.octane.gg/teams/BS+COMPETITION.png"}}},"stat":8041}]

# Group Player Stats

## Player Stats [/stats/players]

### Get Player Stats [GET]

+ Parameters
    + stat: `score` (string, required) - stat names

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"stats":[{"player":{"_id":"60074cee25dd47da00f082fc","slug":"82fc-gordito-bandito","tag":"Gordito Bandito"},"events":[{"_id":"5f35882d53fbbb5894b43037","slug":"3037-rlcs-season-1-north-america-stage-1","name":"RLCS Season 1 North America Stage 1","region":"NA","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs1","rlcsna","rlcs19","rlcs19lp"]}],"teams":[{"_id":"6020be5bf1e4807cc7013208","slug":"3208-elysian-esports","name":"Elysian eSports"}],"opponents":[{"_id":"6020be5bf1e4807cc7013206","slug":"3206-october-sky","name":"October Sky"}],"startDate":"2016-04-30T00:00:00Z","endDate":"2016-04-30T00:00:00Z","games":{"total":3,"replays":0,"wins":2},"matches":{"total":1,"replays":1,"wins":1},"stats":{"score":595}}]}

## Player Stats By Team [/stats/players/teams]

### Get Player Stats By Team [GET]

+ Parameters
    + stat: `score` (string, required) - stat names

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"stats":[{"player":{"_id":"60074cee25dd47da00f082fc","slug":"82fc-gordito-bandito","tag":"Gordito Bandito"},"events":[{"_id":"5f35882d53fbbb5894b43037","slug":"3037-rlcs-season-1-north-america-stage-1","name":"RLCS Season 1 North America Stage 1","region":"NA","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs1","rlcsna","rlcs19","rlcs19lp"]}],"teams":[{"_id":"6020be5bf1e4807cc7013208","slug":"3208-elysian-esports","name":"Elysian eSports"}],"opponents":[{"_id":"6020be5bf1e4807cc7013206","slug":"3206-october-sky","name":"October Sky"}],"startDate":"2016-04-30T00:00:00Z","endDate":"2016-04-30T00:00:00Z","games":{"total":3,"replays":0,"wins":2},"matches":{"total":1,"replays":1,"wins":1},"stats":{"score":595}}]}


## Player Stats By Opponent [/stats/players/opponents]

### Get Player Stats By Opponent [GET]

+ Parameters
    + stat: `score` (string, required) - stat names

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"stats":[{"player":{"_id":"60074cee25dd47da00f082fc","slug":"82fc-gordito-bandito","tag":"Gordito Bandito"},"events":[{"_id":"5f35882d53fbbb5894b43037","slug":"3037-rlcs-season-1-north-america-stage-1","name":"RLCS Season 1 North America Stage 1","region":"NA","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs1","rlcsna","rlcs19","rlcs19lp"]}],"teams":[{"_id":"6020be5bf1e4807cc7013208","slug":"3208-elysian-esports","name":"Elysian eSports"}],"opponents":[{"_id":"6020be5bf1e4807cc7013206","slug":"3206-october-sky","name":"October Sky"}],"startDate":"2016-04-30T00:00:00Z","endDate":"2016-04-30T00:00:00Z","games":{"total":3,"replays":0,"wins":2},"matches":{"total":1,"replays":1,"wins":1},"stats":{"score":595}}]}



## Player Stats By Event [/stats/players/events]

### Get Player Stats By Event [GET]

+ Parameters
    + stat: `score` (string, required) - stat names

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + player: 5f3d8fdd95f40596eae23d97 (string, optional) - a player id

    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"stats":[{"player":{"_id":"60074cee25dd47da00f082fc","slug":"82fc-gordito-bandito","tag":"Gordito Bandito"},"events":[{"_id":"5f35882d53fbbb5894b43037","slug":"3037-rlcs-season-1-north-america-stage-1","name":"RLCS Season 1 North America Stage 1","region":"NA","mode":3,"tier":"S","image":"https://griffon.octane.gg/events/rlcs.png","groups":["rlcs","rlcs1","rlcsna","rlcs19","rlcs19lp"]}],"teams":[{"_id":"6020be5bf1e4807cc7013208","slug":"3208-elysian-esports","name":"Elysian eSports"}],"opponents":[{"_id":"6020be5bf1e4807cc7013206","slug":"3206-october-sky","name":"October Sky"}],"startDate":"2016-04-30T00:00:00Z","endDate":"2016-04-30T00:00:00Z","games":{"total":3,"replays":0,"wins":2},"matches":{"total":1,"replays":1,"wins":1},"stats":{"score":595}}]}

# Group Team Stats

## Team Stats [/stats/teams]

### Get Team Stats [GET]

+ Parameters
    + stat: `score` (string, required) - stat names

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"stats":[{"team":{"_id":"6020bcf2f1e4807cc7007899","slug":"7899-trash-cans","name":"Trash Cans","image":"https://griffon.octane.gg/teams/trash-cans.png"},"events":[{"_id":"5f35882d53fbbb5894b43056","slug":"3056-universal-open-season-1","name":"Universal Open Season 1","region":"INT","mode":2,"tier":"A","image":"https://griffon.octane.gg/events/universal-open.png"}],"players":[{"_id":"5f3d8fdd95f40596eae23e97","slug":"3e97-steazzy","tag":"steazzy"},{"_id":"5f3d8fdd95f40596eae23e94","slug":"3e94-jade","tag":"Jade"}],"opponents":[{"_id":"6020bcf2f1e4807cc7007895","slug":"7895-ludwig-clan","name":"Ludwig Clan","image":"https://griffon.octane.gg/teams/ludwig-clan.png"}],"startDate":"2017-08-06T00:00:00Z","endDate":"2017-08-06T00:00:00Z","games":{"total":2,"replays":0,"wins":0},"matches":{"total":1,"replays":1,"wins":1},"stats":{"score":820}}]}


## Team Stats By Opponent [/stats/teams/opponents]

### Get Team Stats By Opponent [GET]

+ Parameters
    + stat: `score` (string, required) - stat names

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"stats":[{"team":{"_id":"6020bcf2f1e4807cc7007899","slug":"7899-trash-cans","name":"Trash Cans","image":"https://griffon.octane.gg/teams/trash-cans.png"},"events":[{"_id":"5f35882d53fbbb5894b43056","slug":"3056-universal-open-season-1","name":"Universal Open Season 1","region":"INT","mode":2,"tier":"A","image":"https://griffon.octane.gg/events/universal-open.png"}],"players":[{"_id":"5f3d8fdd95f40596eae23e97","slug":"3e97-steazzy","tag":"steazzy"},{"_id":"5f3d8fdd95f40596eae23e94","slug":"3e94-jade","tag":"Jade"}],"opponents":[{"_id":"6020bcf2f1e4807cc7007895","slug":"7895-ludwig-clan","name":"Ludwig Clan","image":"https://griffon.octane.gg/teams/ludwig-clan.png"}],"startDate":"2017-08-06T00:00:00Z","endDate":"2017-08-06T00:00:00Z","games":{"total":2,"replays":0,"wins":0},"matches":{"total":1,"replays":1,"wins":1},"stats":{"score":820}}]}


## Team Stats By Event [/stats/teams/events]

### Get Team Stats By Event [GET]

+ Parameters
    + stat: `score` (string, required) - stat names

    + event: 5f35882d53fbbb5894b43040 (string, optional) - an event id

    + stage: 1 (number, optional) - a stage id

    + match: 6043152fa09e7fba40d2ae62 (string, optional) - a match id

    + qualifier: true (boolean, optional) - stage is a qualifier

    + winner: true (boolean, optional) - game or series winnner

    + nationality: us (string, optional) - a 2-letter country code
    
    + tier: `S` (enum[string], optional) - an event tier  
        + Members
          + S
          + A
          + B
          + C
          + D
          + Monthly
          + Weekly
          + Show Match
          + Qualifier
    
    + region: `NA` (enum[string], optional) - an event region
        + Members
          + NA
          + EU
          + OCE
          + SAM
          + ASIA
          + ME
    
    + mode: `3` (enum[number], optional) - an event mode
        + Members
          + 3
          + 2
          + 1
    
    + group: rlcsx (string, optional) - an event group
    
    + before: `2016-12-03` (date, optional) - filter matches before this date
    
    + after: `2016-12-03` (date, optional) - filter matches after this date

    + bestOf: 5 (enum[number], optional) - a match format
        + Members
          + 3
          + 5
          + 7
    
    + team: 6020bc70f1e4807cc70023c7 (string, optional) - a team id

    + opponent 6020bc70f1e4807cc70023a0 (string, optional) - an opponent team id

* Response 200 (application/json)
  {"stats":[{"team":{"_id":"6020bcf2f1e4807cc7007899","slug":"7899-trash-cans","name":"Trash Cans","image":"https://griffon.octane.gg/teams/trash-cans.png"},"events":[{"_id":"5f35882d53fbbb5894b43056","slug":"3056-universal-open-season-1","name":"Universal Open Season 1","region":"INT","mode":2,"tier":"A","image":"https://griffon.octane.gg/events/universal-open.png"}],"players":[{"_id":"5f3d8fdd95f40596eae23e97","slug":"3e97-steazzy","tag":"steazzy"},{"_id":"5f3d8fdd95f40596eae23e94","slug":"3e94-jade","tag":"Jade"}],"opponents":[{"_id":"6020bcf2f1e4807cc7007895","slug":"7895-ludwig-clan","name":"Ludwig Clan","image":"https://griffon.octane.gg/teams/ludwig-clan.png"}],"startDate":"2017-08-06T00:00:00Z","endDate":"2017-08-06T00:00:00Z","games":{"total":2,"replays":0,"wins":0},"matches":{"total":1,"replays":1,"wins":1},"stats":{"score":820}}]}