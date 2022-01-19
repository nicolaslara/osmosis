(window.webpackJsonp=window.webpackJsonp||[]).push([[54],{504:function(a,s,t){"use strict";t.r(s);var e=t(8),n=Object(e.a)({},(function(){var a=this,s=a.$createElement,t=a._self._c||s;return t("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[t("h1",{attrs:{id:"validating-on-mainnet"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#validating-on-mainnet"}},[a._v("#")]),a._v(" Validating On Mainnet")]),a._v(" "),t("h2",{attrs:{id:"synced-node"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#synced-node"}},[a._v("#")]),a._v(" Synced Node")]),a._v(" "),t("p",[a._v("Before creating a mainnet validator, ensure you have first followed the instructions on how to "),t("a",{attrs:{href:"../developing/network/join-mainnet"}},[a._v("join the mainnet")])]),a._v(" "),t("h2",{attrs:{id:"initialize-wallet-keyring"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#initialize-wallet-keyring"}},[a._v("#")]),a._v(" Initialize Wallet Keyring")]),a._v(" "),t("p",[a._v("If you decide you want to turn your node into a validator, you will first need to add a wallet to your keyring.")]),a._v(" "),t("p",[a._v("While you can add an existing wallet through your seed phrase, we will create a new wallet in this example (replace KEY_NAME with a name of your choosing):")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd keys "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("add")]),a._v(" KEY_NAME\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("Ensure you write down the mnemonic as you can not recover the wallet without it. To ensure your wallet was saved to your keyring, the WALLET_NAME is in your keys list:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd keys list\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("h2",{attrs:{id:"validator-public-key"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#validator-public-key"}},[a._v("#")]),a._v(" Validator Public Key")]),a._v(" "),t("p",[a._v("The last thing needed before initializing the validator is to obtain your validator public key which was created when you first initialized your node. To obtain your validator pubkey:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd tendermint show-validator\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("h2",{attrs:{id:"create-validator-command"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#create-validator-command"}},[a._v("#")]),a._v(" Create Validator Command")]),a._v(" "),t("p",[a._v("Ensure you have a small amount of OSMO on the wallet address you are using on your keyring in order to successfully send a transaction. Once you have have a balance on the address on your keyring, you can now send the create-validator transaction.")]),a._v(" "),t("p",[a._v("Here is the empty command:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd tx staking create-validator "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("KEY_NAME"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--amount"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("staking_amount_uosmo"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--pubkey"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token variable"}},[t("span",{pre:!0,attrs:{class:"token variable"}},[a._v("$(")]),a._v("osmosisd tendermint show-validator"),t("span",{pre:!0,attrs:{class:"token variable"}},[a._v(")")])]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--moniker"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[moniker_id_of_your_node]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--security-contact"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[security contact email/contact method]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[chain-id]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--commission-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[commission_rate]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--commission-max-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[maximum_commission_rate]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--commission-max-change-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[maximum_rate_of_change_of_commission]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--min-self-delegation"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[min_self_delegation_amount]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--gas"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"auto"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--gas-prices"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"[gas_price]"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br"),t("span",{staticClass:"line-number"},[a._v("4")]),t("br"),t("span",{staticClass:"line-number"},[a._v("5")]),t("br"),t("span",{staticClass:"line-number"},[a._v("6")]),t("br"),t("span",{staticClass:"line-number"},[a._v("7")]),t("br"),t("span",{staticClass:"line-number"},[a._v("8")]),t("br"),t("span",{staticClass:"line-number"},[a._v("9")]),t("br"),t("span",{staticClass:"line-number"},[a._v("10")]),t("br"),t("span",{staticClass:"line-number"},[a._v("11")]),t("br"),t("span",{staticClass:"line-number"},[a._v("12")]),t("br"),t("span",{staticClass:"line-number"},[a._v("13")]),t("br")])]),t("p",[a._v("Here is the same command but with example values:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd tx staking create-validator "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("wallet1 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--amount"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("500000000uosmo "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--pubkey"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("osmovalconspub1zcjduepqrevtrgcntyz04w9yzwvpy2ddf2h5pyu2tczgf9dssmywty0tzqzs0gwu0r  "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--moniker"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"Wosmongton"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--security-contact"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"wosmongton@osmosis.labs"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"osmosis-1"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--commission-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.1"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--commission-max-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.2"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--commission-max-change-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.05"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--min-self-delegation"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"500000000"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--gas"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"auto"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n--gas-prices"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.0025uosmo"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br"),t("span",{staticClass:"line-number"},[a._v("4")]),t("br"),t("span",{staticClass:"line-number"},[a._v("5")]),t("br"),t("span",{staticClass:"line-number"},[a._v("6")]),t("br"),t("span",{staticClass:"line-number"},[a._v("7")]),t("br"),t("span",{staticClass:"line-number"},[a._v("8")]),t("br"),t("span",{staticClass:"line-number"},[a._v("9")]),t("br"),t("span",{staticClass:"line-number"},[a._v("10")]),t("br"),t("span",{staticClass:"line-number"},[a._v("11")]),t("br"),t("span",{staticClass:"line-number"},[a._v("12")]),t("br"),t("span",{staticClass:"line-number"},[a._v("13")]),t("br")])]),t("p",[a._v("If you need further explanation for each of these command flags:")]),a._v(" "),t("ul",[t("li",[a._v("the "),t("code",[a._v("from")]),a._v(" flag is the KEY_NAME you created when initializing the key on your keyring")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("amount")]),a._v(" flag is the amount you will place in your own validator in uosmo (in the example, 500000000uosmo is 500osmo)")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("pubkey")]),a._v(" is the validator public key found earlier")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("moniker")]),a._v(" is a human readable name you choose for your validator")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("security-contact")]),a._v(" is an email your delegates are able to contact you at")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("chain-id")]),a._v(" is whatever chain-id you are working with (in the osmosis mainnet case it is osmosis-1)")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("commission-rate")]),a._v(" is the rate you will charge your delegates (in the example above, 10 percent)")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("commission-max-rate")]),a._v(" is the most you are allowed to charge your delegates (in the example above, 20 percent)")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("commission-max-change-rate")]),a._v(" is how much you can increase your commission rate in a 24 hour period (in the example above, 5 percent per day until reaching the max rate)")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("min-self-delegation")]),a._v(" is the lowest amount of personal funds the validator is required to have in their own validator to stay bonded (in the example above, 500osmo)")]),a._v(" "),t("li",[a._v("the "),t("code",[a._v("gas-prices")]),a._v(" is the amount of gas used to send this create-validator transaction")])]),a._v(" "),t("h2",{attrs:{id:"track-validator-active-set"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#track-validator-active-set"}},[a._v("#")]),a._v(" Track Validator Active Set")]),a._v(" "),t("p",[a._v("To see the current validator active set:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd query staking validators -o json "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("|")]),a._v(" jq -r "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v("'.validators[] | \n[.operator_address, .status, (.tokens|tonumber / pow(10; 6)), \n.commission.update_time[0:19], .description.moniker] | @csv'")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("|")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("column")]),a._v(" -t -s"),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('","')]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br")])]),t("p",[a._v("You can search for your specific moniker by adding grep MONIKER at the end:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd query staking validators -o json "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("|")]),a._v(" jq -r "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v("'.validators[] | \n[.operator_address, .status, (.tokens|tonumber / pow(10; 6)), \n.commission.update_time[0:19], .description.moniker] | @csv'")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("|")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("column")]),a._v(" -t -s"),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('","')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("|")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("grep")]),a._v(" Wosmongton\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br")])]),t("p",[a._v("If your bond status is "),t("code",[a._v("BOND_STATUS_BONDED")]),a._v(", congratulations, your validator is part of the active validator set!")]),a._v(" "),t("p",[a._v("Please note, as of this writing, you must be in the top 100 validators (in other words, must have more OSMO delegated to your validator than the 100th validator in the active validator set) to be bonded. If you did everything above correct but do not have more OSMO delegated to your validator than the 100th validator, you will stay unbonded.")]),a._v(" "),t("h2",{attrs:{id:"track-validator-signing"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#track-validator-signing"}},[a._v("#")]),a._v(" Track Validator Signing")]),a._v(" "),t("p",[a._v("To track your validator's signing history, copy the validator public key:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd tendermint show-validator\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("Use your validators public key queried above as the validator-pubkey below:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd query slashing signing-info "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),a._v("validator-pubkey"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("Example:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("osmosisd query slashing signing-info "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('\'{"@type":"/cosmos.crypto.ed25519.PubKey","key":"HlixoxNZBPq4pBOYEimtSq9Ak4peBISVsIbI5ZHrEAU="}\'')]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])])])}),[],!1,null,null,null);s.default=n.exports}}]);