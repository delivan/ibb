<template>
	<div class="user-deposit">
		<div class="deposit-status">
			<div class="deposit-balance">
				<div class="title">Deposit Balance</div>
				<div class="value">${{ pools.reduce((acc, pool) => (acc += ((pool.AssetDeposit / 1000000) * pool.AssetPrice) / 1000000), 0).toFixed(2) }}</div>
			</div>
			<div class="net-apy">
				<div class="title">Net APY</div>
				<div class="value">{{ pools.reduce((acc, pool) => (acc += pool.AssetDeposit ? pool.DepositApy / 10000 : 0), 0).toFixed(2) }}%</div>
			</div>
		</div>
		<div class="asset-table">
			<div class="table-header">
				<div class="table-cell"><span>Asset</span></div>
				<div class="table-cell"><span>APY</span></div>
				<div class="table-cell"><span>Balance</span></div>
				<!-- <div class="table-cell"><span>Collateral</span></div> -->
				<div class="table-cell"><span>Claim Interest</span></div>
			</div>
			<div v-if="Array.isArray(pools) && pools.length > 0" class="table-rows">
				<div v-for="(pool, index) in pools" v-bind:key="pool.id" class="table-row" @click="clickAsset(pool)">
					<div class="table-cell">{{ pool.Asset }}</div>
					<div class="table-cell">{{ parseFloat(pool.DepositApy / 10000).toFixed(2) }}%</div>
					<div class="table-cell">{{ pool.AssetDeposit / 1000000 }}</div>
					<!-- <div class="table-cell"><input type="checkbox" v-model="pool.Collatoral" /></div> -->
					<div class="table-cell">
						<button v-if="loadingClaimPoolIndecies.some((loadingIndex) => loadingIndex === index)" @click="clickClaim(pool, $event, index)">Loading...</button>
						<button v-else @click="clickClaim(pool, $event, index)">Claim {{ pool.DepositEarned / 1000000 }}</button>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.user-deposit {
	width: 100%;
	padding: 0 20px;
}

.deposit-status {
	display: flex;
	justify-content: space-between;
}

.title {
	font-size: 20px;
}

.value {
	font-size: 26px;
	font-weight: bold;
}

.asset-table {
	margin-top: 16px;
	padding: 12px;
	border: 1px solid white;
	border-radius: 10px;
}

.table-header {
	display: flex;
	padding: 8px 12px;
	color: rgba(255, 255, 255, 0.7);
	font-size: 14px;
}

.table-rows {
	font-weight: bold;
}

.table-row {
	display: flex;
	padding: 8px 12px 7px;
	border-radius: 5px;
}

.table-row:hover {
	cursor: pointer;
	background: rgba(0, 0, 0, 0.5);
}

.table-cell {
	width: 100%;
	min-width: 40px;
	display: flex;
	align-items: center;
	justify-content: flex-end;
}

.table-cell:first-child {
	width: 10px;
	justify-content: flex-start;
}

.table-cell > span {
	line-height: 14px;
}

.table-cell > button {
	background: transparent;
	border: 1px solid white;
	border-radius: 3px;
	color: white;
	padding: 2px 4px;
}

.table-cell > button:hover {
	cursor: pointer;
}
</style>

<script>
export default {
	name: 'UserDeposit',
	data() {
		return {
			loadingClaimPoolIndecies: []
		}
	},
	computed: {
		pools() {
			const loggedAddress = this.$store.getters['common/wallet/address']
			const userAssets = loggedAddress
				? this.$store.getters['sapienscosmos.ibb.ibb/getUserLoad']({
						params: {
							id: loggedAddress
						}
				  })?.LoadUserResponse ?? []
				: []
			const assetPools =
				this.$store.getters['sapienscosmos.ibb.ibb/getPoolLoad']({
					params: {}
				})?.LoadPoolResponse ?? []

			return assetPools
				.map((pool, index) => ({
					...pool,
					...userAssets[index]
				}))
				.filter((pool) => pool.AssetDeposit > 0)
		}
	},
	methods: {
		clickAsset(pool) {
			this.$emit('click-asset', pool, 'Deposit')
		},
		async clickClaim(pool, e, loadingIndex) {
			e.stopPropagation()
			this.loadingClaimPoolIndecies = this.loadingClaimPoolIndecies.concat([loadingIndex])
			const value = {
				creator: this.$store.getters['common/wallet/address'],
				asset: pool.Asset,
				denom: `u${pool.Asset.toLocaleLowerCase()}`
			}
			await this.$store.dispatch(`sapienscosmos.ibb.ibb/sendMsgCreateClaim2`, {
				value,
				fee: []
			})
			await this.$store.dispatch('sapienscosmos.ibb.ibb/QueryUserLoad', {
				options: { all: true },
				params: {
					id: this.$store.getters['common/wallet/address']
				}
			})
			this.loadingClaimPoolIndecies = this.loadingClaimPoolIndecies.filter((index) => index !== loadingIndex)
		}
	}
}
</script>
