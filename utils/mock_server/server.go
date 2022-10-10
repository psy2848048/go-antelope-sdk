package mockServer

import (
	"github.com/jarcoal/httpmock"
	"github.com/psy2848048/go-antelope-sdk/utils"
)

func CreateAndActivateRestMockServer() {
	httpmock.Activate()

	httpmock.RegisterResponder(
		utils.POST, "http://localhost/v1/chain/get_producers",
		httpmock.NewStringResponder(200, chainAPIGetProducersResp),
	)
}

func DeactivateMockServer() {
	httpmock.DeactivateAndReset()
}

const chainAPIGetProducersResp = `
{
    "rows": [
        {
            "owner": "eosflytomars",
            "total_votes": "12364922069802956800.00000000000000000",
            "producer_key": "EOS6Agpfp38bTyRjJDmB4Qb1EpQSq7wnEAsALXgXE7KFSzKjokkFD",
            "is_active": 1,
            "url": "https://www.bitmars.one",
            "unpaid_blocks": 2523,
            "last_claim_time": "2022-10-09T01:22:54.000",
            "location": 156
        },
        {
            "owner": "binancestake",
            "total_votes": "12352903870702721024.00000000000000000",
            "producer_key": "EOS7unwwVJfmKonrT6Gj46LDiNUPpFhpPALpTe2eofmFeoG74bKKn",
            "is_active": 1,
            "url": "https://binance-eos.com",
            "unpaid_blocks": 4308,
            "last_claim_time": "2022-10-08T20:11:00.000",
            "location": 136
        },
        {
            "owner": "newdex.bp",
            "total_votes": "12347824646442350592.00000000000000000",
            "producer_key": "EOS688SnH8tQ7NiyhamiCzWXAGPDLF9S7K8ga79UBHKFgjS1MhqhB",
            "is_active": 1,
            "url": "https://newpool.io",
            "unpaid_blocks": 2988,
            "last_claim_time": "2022-10-09T00:02:55.500",
            "location": 392
        },
        {
            "owner": "eosnationftw",
            "total_votes": "12272880139535597568.00000000000000000",
            "producer_key": "EOS8L12yBrtx7mpewHmjwgJeNb2aLaeQdoDgMW82dzDSu17ec2XNL",
            "is_active": 1,
            "url": "https://eosnation.io",
            "unpaid_blocks": 2138,
            "last_claim_time": "2022-10-09T02:30:32.000",
            "location": 124
        },
        {
            "owner": "eosinfstones",
            "total_votes": "12208759554479728640.00000000000000000",
            "producer_key": "EOS6CSvGzNhNxVYbcnWSuheNcfzjGeGBY9trR4YAJ4Yvakq4oCh6y",
            "is_active": 1,
            "url": "http://infinitystones.io",
            "unpaid_blocks": 7526,
            "last_claim_time": "2022-10-08T10:46:00.000",
            "location": 1
        },
        {
            "owner": "eoscannonchn",
            "total_votes": "11937099305900587008.00000000000000000",
            "producer_key": "EOS73cTi9V7PNg4ujW5QzoTfRSdhH44MPiUJkUV6m3oGwj7RX7kML",
            "is_active": 1,
            "url": "https://eoscannon.io",
            "unpaid_blocks": 1956,
            "last_claim_time": "2022-10-09T03:02:10.500",
            "location": 136,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS73cTi9V7PNg4ujW5QzoTfRSdhH44MPiUJkUV6m3oGwj7RX7kML",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "aus1genereos",
            "total_votes": "11659262796010577920.00000000000000000",
            "producer_key": "EOS7rJvrhEAs9zqXewR9iyhsFNsarbu3EoDJRkYTHz2b87icsxhuA",
            "is_active": 1,
            "url": "https://genereos.io",
            "unpaid_blocks": 2986,
            "last_claim_time": "2022-10-08T23:59:36.000",
            "location": 36,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS7rJvrhEAs9zqXewR9iyhsFNsarbu3EoDJRkYTHz2b87icsxhuA",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "hashfineosio",
            "total_votes": "11622602410212988928.00000000000000000",
            "producer_key": "EOS7jSfvStvbKDmGvQdtrQsCyNkWczXfvh6CHmBVmeypJyHsUrMqj",
            "is_active": 1,
            "url": "https://www.hashfin.com",
            "unpaid_blocks": 1644,
            "last_claim_time": "2022-10-09T03:59:18.000",
            "location": 0,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS7jSfvStvbKDmGvQdtrQsCyNkWczXfvh6CHmBVmeypJyHsUrMqj",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "big.one",
            "total_votes": "11538600873294366720.00000000000000000",
            "producer_key": "EOS8MpYyXwn3DLqk9Y9XTHYcd6wGGijNqJefFoQEwEoXTq1awZ42w",
            "is_active": 1,
            "url": "https://eos.big.one",
            "unpaid_blocks": 1044,
            "last_claim_time": "2022-10-09T05:41:39.500",
            "location": 392
        },
        {
            "owner": "okcapitalbp1",
            "total_votes": "11038934828171313152.00000000000000000",
            "producer_key": "EOS6NqWZ1i9KSNoeBiby6Nmf1seAbEfhvrDoCbwSi1hV4cuqqnYRP",
            "is_active": 1,
            "url": "https://www.okex.com/eosbp/",
            "unpaid_blocks": 7184,
            "last_claim_time": "2022-10-08T11:46:09.500",
            "location": 0
        },
        {
            "owner": "starteosiobp",
            "total_votes": "10884584491019692032.00000000000000000",
            "producer_key": "EOS4wZZXm994byKANLuwHD6tV3R3Mu3ktc41aSVXCBaGnXJZJ4pwF",
            "is_active": 1,
            "url": "https://www.starteos.io",
            "unpaid_blocks": 2835,
            "last_claim_time": "2022-10-09T00:26:54.000",
            "location": 156
        },
        {
            "owner": "atticlabeosb",
            "total_votes": "10844864432195651584.00000000000000000",
            "producer_key": "EOS7PfA3A4UdfMu2wKbuXdbHn8EWAxbMnFoFWui4X2zsr2oPwdQJP",
            "is_active": 1,
            "url": "https://atticlab.net",
            "unpaid_blocks": 5253,
            "last_claim_time": "2022-10-08T17:24:06.000",
            "location": 804
        },
        {
            "owner": "eoslaomaocom",
            "total_votes": "10440648946091925504.00000000000000000",
            "producer_key": "EOS8QgURqo875qu3a8vgZ58qBeu2cTehe9zAWRfpdCXAQipicu1Fi",
            "is_active": 1,
            "url": "https://eoslaomao.com",
            "unpaid_blocks": 10512,
            "last_claim_time": "2022-10-08T02:05:49.500",
            "location": 392,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS8QgURqo875qu3a8vgZ58qBeu2cTehe9zAWRfpdCXAQipicu1Fi",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "blockpooleos",
            "total_votes": "10090362309438683136.00000000000000000",
            "producer_key": "EOS61FDJz3GC42GhaPSsmKh7SxuesyZhjm7hBwBKqN52v1HukEqBu",
            "is_active": 1,
            "url": "https://www.blockpool.com",
            "unpaid_blocks": 1248,
            "last_claim_time": "2022-10-09T05:05:42.000",
            "location": 0
        },
        {
            "owner": "eosiosg11111",
            "total_votes": "10072919871277869056.00000000000000000",
            "producer_key": "EOS7zVBQMhV7dZ5zRQwBgDmmbFCHA6YcmwW6Dq5CePGpqLR1ZsVAc",
            "is_active": 1,
            "url": "https://eosio.sg",
            "unpaid_blocks": 5844,
            "last_claim_time": "2022-10-08T15:43:36.000",
            "location": 702
        },
        {
            "owner": "bitfinexeos1",
            "total_votes": "10063513417383008256.00000000000000000",
            "producer_key": "EOS4tkw7LgtURT3dvG3kQ4D1sg3aAtPDymmoatpuFkQMc7wzZdKxc",
            "is_active": 1,
            "url": "https://www.bitfinex.com",
            "unpaid_blocks": 2040,
            "last_claim_time": "2022-10-09T02:48:20.000",
            "location": 0
        },
        {
            "owner": "teamgreymass",
            "total_votes": "9950096137996158976.00000000000000000",
            "producer_key": "EOS5ktvwSdLEdusdRn7NmdV2Xu89xiXjir7EhJuZ4DUa8WMNuojbx",
            "is_active": 1,
            "url": "https://greymass.com",
            "unpaid_blocks": 6215,
            "last_claim_time": "2022-10-08T14:38:02.500",
            "location": 124,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS5ktvwSdLEdusdRn7NmdV2Xu89xiXjir7EhJuZ4DUa8WMNuojbx",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "whaleex.com",
            "total_votes": "9933500342574950400.00000000000000000",
            "producer_key": "EOS88EGcFghfQJER1mDaEe4kDJ7MGDoPmXQfA7q2QMTLLqiYP1UQR",
            "is_active": 1,
            "url": "https://www.whaleex.com",
            "unpaid_blocks": 6148,
            "last_claim_time": "2022-10-08T14:48:16.000",
            "location": 410
        },
        {
            "owner": "eosasia11111",
            "total_votes": "9922242178808612864.00000000000000000",
            "producer_key": "EOS76gG6ATpqfVf5KrVjh3f4JAa4EKzAwWabTucNQ4Xv2TmVAj9bN",
            "is_active": 1,
            "url": "https://www.eosasia.one/",
            "unpaid_blocks": 6396,
            "last_claim_time": "2022-10-08T14:06:09.500",
            "location": 9
        },
        {
            "owner": "ivote4eosusa",
            "total_votes": "9906453402474604544.00000000000000000",
            "producer_key": "EOS6KzD4YVbuXV5uBH5d4Ay4sTzuQk88ivmnWfJPLoo6SFrX6iyqj",
            "is_active": 1,
            "url": "https://eosusa.io/bp/eos/",
            "unpaid_blocks": 2268,
            "last_claim_time": "2022-10-09T02:10:06.000",
            "location": 840,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS6KzD4YVbuXV5uBH5d4Ay4sTzuQk88ivmnWfJPLoo6SFrX6iyqj",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "eosrapidprod",
            "total_votes": "9197784592100426752.00000000000000000",
            "producer_key": "EOS8QEFsgUWj7BscQNkiremtpSoRkzwDqmCPpKKCHYXGNaqxXFQ4h",
            "is_active": 1,
            "url": "https://eosrapid.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T19:12:00.000",
            "location": 840
        },
        {
            "owner": "eoslambdacom",
            "total_votes": "8384089495966818304.00000000000000000",
            "producer_key": "EOS66Guagdf4FSN9diXivjeW5tV8yPqBrQFFusTaN7rLTWD5LsnSN",
            "is_active": 1,
            "url": "https://bridgepool.lambda.im",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T06:17:07.000",
            "location": 0,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS66Guagdf4FSN9diXivjeW5tV8yPqBrQFFusTaN7rLTWD5LsnSN",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "eosphereiobp",
            "total_votes": "8156807138283275264.00000000000000000",
            "producer_key": "EOS5FYRrG3ThNU56d3uZ8d4bo4MsN5Fig9WPY3xYSpH3RqPhpoH3A",
            "is_active": 1,
            "url": "https://www.eosphere.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T15:43:03.000",
            "location": 36
        },
        {
            "owner": "truststaking",
            "total_votes": "8042369827713352704.00000000000000000",
            "producer_key": "EOS5jZgX1gZvYat588FKsMyAybVo8tWaAX4bHdBrW9Z4C5iMshUKf",
            "is_active": 1,
            "url": "https://www.binance.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T20:11:07.000",
            "location": 392
        },
        {
            "owner": "eosdotwikibp",
            "total_votes": "6796798574209761280.00000000000000000",
            "producer_key": "EOS7RsdDs8k8GDAdZrETnTjoGwiqAwwdNyxeH8q6fmHgpHjPPnyco",
            "is_active": 1,
            "url": "https://eos.wiki",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T07:33:12.000",
            "location": 276,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS7RsdDs8k8GDAdZrETnTjoGwiqAwwdNyxeH8q6fmHgpHjPPnyco",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "eosx.game",
            "total_votes": "6136821275287072768.00000000000000000",
            "producer_key": "EOS727yFoVQYBsXZ9qxtZJPJu3rm4befuoVwBk7iqgBWjmEGAMUax",
            "is_active": 1,
            "url": "https://chaingame.club",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T03:02:00.500",
            "location": 156
        },
        {
            "owner": "moreisfuture",
            "total_votes": "5622908315041492992.00000000000000000",
            "producer_key": "EOS671RDSn2embqDJhFHKWZqqCfVgNAmi2p3vQZfQBX77GypcGoF2",
            "is_active": 1,
            "url": "https://more.top",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T04:56:48.000",
            "location": 156
        },
        {
            "owner": "slowmistiobp",
            "total_votes": "5533449115910307840.00000000000000000",
            "producer_key": "EOS62hJuwYdARgXDWLepwTP4Xc749Gi3RVXZp1qJ8NjfbEwi79uZN",
            "is_active": 1,
            "url": "https://slowmist.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T08:20:13.000",
            "location": 0
        },
        {
            "owner": "bp.defi",
            "total_votes": "5266186696687296512.00000000000000000",
            "producer_key": "EOS5BoXgRJwL7JFvKnV64Q3Ha3ux6x2cP8nnhU9NVrRkyrhPC3m5b",
            "is_active": 1,
            "url": "https://defibox.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T09:29:02.500",
            "location": 392,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS5BoXgRJwL7JFvKnV64Q3Ha3ux6x2cP8nnhU9NVrRkyrhPC3m5b",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "eosauthority",
            "total_votes": "4976294023303766016.00000000000000000",
            "producer_key": "EOS4va3CTmAcAAXsT26T3EBWqYHgQLshyxsozYRgxWm9tjmy17pVV",
            "is_active": 1,
            "url": "https://eosauthority.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T04:05:51.500",
            "location": 826
        },
        {
            "owner": "zbeosbp11111",
            "total_votes": "4824425275959604224.00000000000000000",
            "producer_key": "EOS7rhgVPWWyfMqjSbNdndtCK8Gkza3xnDbUupsPLMZ6gjfQ4nX81",
            "is_active": 1,
            "url": "https://www.zbeos.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-08-18T05:01:48.000",
            "location": 156
        },
        {
            "owner": "eos42freedom",
            "total_votes": "4692402051407156224.00000000000000000",
            "producer_key": "EOS4tw7vH62TcVtMgm2tjXzn9hTuHEBbGPUK2eos42ssY7ip4LTzu",
            "is_active": 1,
            "url": "https://eos42.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T16:30:12.000",
            "location": 276,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS4tw7vH62TcVtMgm2tjXzn9hTuHEBbGPUK2eos42ssY7ip4LTzu",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "helloeoscnbp",
            "total_votes": "4526907321072028160.00000000000000000",
            "producer_key": "EOS79cHpaEittzgJWgj79tdRhgzLEWy8wXmmQ3fL7kkDjmYYiGNet",
            "is_active": 1,
            "url": "https://www.helloeos.com.cn",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T10:31:12.000",
            "location": 156
        },
        {
            "owner": "eosrainbowbp",
            "total_votes": "4512420268146845184.00000000000000000",
            "producer_key": "EOS62q4Q4mg6GkACM7V5mrC46oZiSmV7roJDFA2N11JBdMLhYjRtt",
            "is_active": 1,
            "url": "eosrainbow.xyz",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T09:11:49.500",
            "location": 0
        },
        {
            "owner": "dexeosbpnode",
            "total_votes": "4494450990952401920.00000000000000000",
            "producer_key": "EOS5qraVVdEwHaouJEqpFWJfZuxQViDxP3WZ826wGwt2bxnfuEGRF",
            "is_active": 1,
            "url": "https://dexeos.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T07:24:48.000",
            "location": 0
        },
        {
            "owner": "bp.pizza",
            "total_votes": "4377532358591153664.00000000000000000",
            "producer_key": "EOS8AZhG9bKGPYmoXGvBioZEJYqVj5xMfujivYnHhwJ3SB9oUWty2",
            "is_active": 1,
            "url": "https://pizza.finance",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T03:00:36.000",
            "location": 156
        },
        {
            "owner": "validatoreos",
            "total_votes": "4261372144849701888.00000000000000000",
            "producer_key": "EOS8h87g5rpLg77ZCJy1J2JPJaCtEeNTAbhGmEsQoz5ngGt3cRaHX",
            "is_active": 1,
            "url": "http://validatoreos.xyz",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T13:41:51.000",
            "location": 0
        },
        {
            "owner": "eostitanprod",
            "total_votes": "4258536806845999616.00000000000000000",
            "producer_key": "EOS5z4JJRLvNasG4VRWYLRNC5TeLquNM8k5oBWWpRhpf9mgJH9e9K",
            "is_active": 1,
            "url": "https://eostitan.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T14:34:44.000",
            "location": 0
        },
        {
            "owner": "ecoboost1111",
            "total_votes": "4230239510400140800.00000000000000000",
            "producer_key": "EOS5g9D7uJuEwNzby2uAu1qVxLishXJm1T2odBPxdU3UzP7xYwVA5",
            "is_active": 1,
            "url": "https://ecoboost.app/",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T09:21:18.000",
            "location": 0
        },
        {
            "owner": "hexlantttttt",
            "total_votes": "4184478077341647872.00000000000000000",
            "producer_key": "EOS83xBpEBigshibgFNjbX3jqHZNDzcnJJrGZaq2ioYTjpGEfckXf",
            "is_active": 1,
            "url": "https://hexbp.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T02:40:24.000",
            "location": 410
        },
        {
            "owner": "alohaeosprod",
            "total_votes": "4158986265906736128.00000000000000000",
            "producer_key": "EOS53pfXfxZ4tH3EGccdnGvBZVJsrcSf2nbCKiLLMphgaii9XxxhM",
            "is_active": 1,
            "url": "https://www.alohaeos.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T03:12:32.500",
            "location": 840
        },
        {
            "owner": "cochainworld",
            "total_votes": "4065629301285115904.00000000000000000",
            "producer_key": "EOS5QDSQyh96SmktA28dteEchc1QCVdqKYG4YVDHGGATMYxqy33wi",
            "is_active": 1,
            "url": "https://eoscochain.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T03:39:12.000",
            "location": 156
        },
        {
            "owner": "argentinaeos",
            "total_votes": "4045112710775138304.00000000000000000",
            "producer_key": "EOS7jq4FHrFrtCXxpRQ39dBeDMa5AjM4VaRbqBECkSa5aZnizJzrx",
            "is_active": 1,
            "url": "https://www.eosargentina.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T15:05:11.500",
            "location": 32
        },
        {
            "owner": "geosoneforbp",
            "total_votes": "4037676176202859520.00000000000000000",
            "producer_key": "EOS7CjPeMYUj2Ydmw667JJNsQA5mige9YE4FVxMCT81Jst3SRVbX3",
            "is_active": 1,
            "url": "https://www.geos.one",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T03:28:29.000",
            "location": 0
        },
        {
            "owner": "eoshenzhenio",
            "total_votes": "4032463633357993984.00000000000000000",
            "producer_key": "EOS8EJrMphgHJx5EkHQ4ryodbvnocZEC1xp5T1FnRsT7XtzYA1FE7",
            "is_active": 1,
            "url": "https://info.eoshenzhen.io:1443",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T16:42:54.000",
            "location": 156
        },
        {
            "owner": "kikifinance1",
            "total_votes": "4016967047516044288.00000000000000000",
            "producer_key": "EOS7CioAtoABvsAbmruLAQdDdPNqSKXka8cp1S1zvWBAk8qvX2AN3",
            "is_active": 1,
            "url": "https://kiki.finance",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T04:47:18.000",
            "location": 0,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS7CioAtoABvsAbmruLAQdDdPNqSKXka8cp1S1zvWBAk8qvX2AN3",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "eosflareiobp",
            "total_votes": "3925556091134896128.00000000000000000",
            "producer_key": "EOS6FLARENUBrQSqSRWk4BZG6oyUcepnTb1T6j3LgrExpJaL4gk89",
            "is_active": 1,
            "url": "https://eosflare.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T05:10:30.000",
            "location": 840
        },
        {
            "owner": "cryptolions1",
            "total_votes": "3863421668723772416.00000000000000000",
            "producer_key": "EOS5yHaUBwhpwFftNjADdVwBpCVybLPg5P3z7HpTbDGjUfpqwWSSf",
            "is_active": 1,
            "url": "https://cryptolions.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T23:52:05.000",
            "location": 804
        },
        {
            "owner": "itokenpocket",
            "total_votes": "3852896259427981824.00000000000000000",
            "producer_key": "EOS8BYCigFER1XBbMqD3GjVVF5B3BmB1E1nkjJJTxPX4tmdS3YwMR",
            "is_active": 1,
            "url": "https://www.tokenpocket.pro",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T08:23:42.000",
            "location": 800
        },
        {
            "owner": "eospaceioeos",
            "total_votes": "3753549309894725632.00000000000000000",
            "producer_key": "EOS7dZPHgFJXwHz84JWf5QgybxLpp95cxYE5c4yngsrSqgVFczaPs",
            "is_active": 1,
            "url": "https://eospacex.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T06:57:30.000",
            "location": 156
        },
        {
            "owner": "eosriobrazil",
            "total_votes": "3658126321891840000.00000000000000000",
            "producer_key": "EOS62VTMkGDMWpK3ECNegVLt92UYQQMXC3uzcx5zjW64hrMCaL1rx",
            "is_active": 1,
            "url": "https://eosrio.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T06:31:22.000",
            "location": 76
        },
        {
            "owner": "eosamsterdam",
            "total_votes": "3627550843465688064.00000000000000000",
            "producer_key": "EOS6jd8kF7KxAv87Q9cAmjcyxuJSTKu47jpbJjFebJhYMMXJoqGU1",
            "is_active": 1,
            "url": "https://eosamsterdam.net",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T10:43:12.000",
            "location": 1,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS6jd8kF7KxAv87Q9cAmjcyxuJSTKu47jpbJjFebJhYMMXJoqGU1",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "eosiodetroit",
            "total_votes": "3545132940723919872.00000000000000000",
            "producer_key": "EOS8XCTx97A3ZBM5KgAZkZMb2k2feYK1ukJzMuEX7pUXW1TMG8spt",
            "is_active": 1,
            "url": "https://detroitledger.tech",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T00:45:06.000",
            "location": 840,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS8XCTx97A3ZBM5KgAZkZMb2k2feYK1ukJzMuEX7pUXW1TMG8spt",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "bp.dfs",
            "total_votes": "3513937325349957632.00000000000000000",
            "producer_key": "EOS7u2ARuKEDN2ybqSNY5S2cuoLZ6jFBT3BjU5g8SdHsU8kyEESso",
            "is_active": 1,
            "url": "https://defis.network",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T22:54:30.000",
            "location": 702
        },
        {
            "owner": "eossv12eossv",
            "total_votes": "3342150493125626880.00000000000000000",
            "producer_key": "EOS7fv6nadePJBdCiWyR2Ldz5YWyzJGMBkEo47zUSWFttduQhHHaa",
            "is_active": 1,
            "url": "https://eossiliconvalley.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T23:21:48.500",
            "location": 840
        },
        {
            "owner": "eosathenabp1",
            "total_votes": "3268538743732166656.00000000000000000",
            "producer_key": "EOS6NVNPZDdmBvqMYdifQ4QueT8fr7wDHyxNMgFEKLx9z1Dt5eFE1",
            "is_active": 1,
            "url": "http://athenbp.club",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T17:09:45.500",
            "location": 0
        },
        {
            "owner": "eosbeijingbp",
            "total_votes": "2942059209233420288.00000000000000000",
            "producer_key": "EOS5dGpcEhwB4VEhhXEcqtZs9EQj5HeetuXDnsAGVHMXHAFdMjbdj",
            "is_active": 1,
            "url": "https://www.eosbeijing.one",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T08:00:30.000",
            "location": 392
        },
        {
            "owner": "eossnzpoolbp",
            "total_votes": "2688186882652805120.00000000000000000",
            "producer_key": "EOS6gQeftj3V2yuGdxAJcNfNAw2PpsjftYK4PrmW6uob6xgYcQHKw",
            "is_active": 1,
            "url": "http://snzholding.com/",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T10:43:26.500",
            "location": 1
        },
        {
            "owner": "eosbixinboot",
            "total_votes": "2673487531063319552.00000000000000000",
            "producer_key": "EOS7QC1XfAtkYeLjbHQjcDWVqUsxuSJ3YRhNyG93VAv2u3uHopGVp",
            "is_active": 1,
            "url": "https://www.eosbixin.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T06:59:14.500",
            "location": 7
        },
        {
            "owner": "eosnodeonebp",
            "total_votes": "2628745688246179840.00000000000000000",
            "producer_key": "EOS6hRwZ76vaDVwqEd9oiD2TPtr5HypxiXmiP34pkuicyFDFd6SLN",
            "is_active": 1,
            "url": "https://www.nodeone.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T21:49:24.000",
            "location": 410,
            "producer_authority": [
                "block_signing_authority_v0",
                {
                    "threshold": 1,
                    "keys": [
                        {
                            "key": "EOS6hRwZ76vaDVwqEd9oiD2TPtr5HypxiXmiP34pkuicyFDFd6SLN",
                            "weight": 1
                        }
                    ]
                }
            ]
        },
        {
            "owner": "certikeosorg",
            "total_votes": "2526714049166454784.00000000000000000",
            "producer_key": "EOS57kRcpx1QT4Wn4uA6agwYdJNGew4cfrzKnzrvQZiVE85huUBAB",
            "is_active": 1,
            "url": "https://certik.org",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T11:11:49.000",
            "location": 840
        },
        {
            "owner": "hoo.com",
            "total_votes": "2407348698478720512.00000000000000000",
            "producer_key": "EOS7Ef2ReXFWYmp1SB1ZzhnhDMfY8NnMZnvo4c51ZxiDf2UE7nCNR",
            "is_active": 1,
            "url": "https://hoo.com",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-07-29T15:59:24.500",
            "location": 0
        },
        {
            "owner": "eosvenezuela",
            "total_votes": "2360673180157740544.00000000000000000",
            "producer_key": "EOS7pUMws312EjWwCcyMvV8hJiN5DFePdxsCknTJsrthp1an2AZXu",
            "is_active": 1,
            "url": "https://eosvenezuela.io",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-08T10:00:04.500",
            "location": 862
        },
        {
            "owner": "eosswedenorg",
            "total_votes": "1949848286880204032.00000000000000000",
            "producer_key": "EOS7SGSBsWhSob6TEric6u3TGodcc1uXFcqSrquJ3Etuqcbb3VnNY",
            "is_active": 1,
            "url": "https://eossweden.org",
            "unpaid_blocks": 0,
            "last_claim_time": "2022-10-09T06:40:42.000",
            "location": 752
        }
    ],
    "total_producer_vote_weight": "440324019838359830528.00000000000000000",
    "more": "eoscafeblock"
}
`