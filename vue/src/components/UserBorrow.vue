<template>
	<div class="user-borrow">
		<div class="borrow-status">
			<div class="borrow-balance">
				<div class="title">Borrow Balance</div>
				<div class="value">
					${{ Array.isArray(pools) ? pools.reduce((acc, pool) => (acc += ((pool.AssetBorrow / 1000000) * pool.AssetPrice) / 1000000), 0).toFixed(2) : 0 }}
				</div>
			</div>
			<div class="net-apy">
				<div class="title">Borrow Limit</div>
				<div class="value">${{ parseFloat(currentBollowLimit).toFixed(2) }}</div>
			</div>
		</div>
		<div class="asset-table">
			<div class="table-header">
				<div class="table-cell"><span>Asset</span></div>
				<div class="table-cell"><span>APY</span></div>
				<div class="table-cell"><span>Balance</span></div>
				<div class="table-cell"><span>Interest</span></div>
			</div>
			<div v-if="Array.isArray(pools)" class="table-rows">
				<div v-for="pool in pools" v-bind:key="pool.id" class="table-row" @click="clickAsset(pool)">
					<div class="table-cell">{{ pool.Asset }}</div>
					<div class="table-cell">{{ parseFloat(pool.BorrowApy / 10000).toFixed(2) }}%</div>
					<div class="table-cell">{{ pool.AssetBorrow / 1000000 }}</div>
					<div class="table-cell">{{ pool.BorrowAccrued / 100000 }}</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.user-borrow {
	width: 100%;
	padding: 0 20px;
}

.borrow-status {
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
</style>

<script>
export default {
	name: 'UserBorrow',
	computed: {
		assetPoolsWithUser() {
			const loggedAddress = this.$store.getters['common/wallet/address']
			if (!loggedAddress) {
				return []
			}

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
			return assetPools.map((pool, index) => ({
				...pool,
				...userAssets[index]
			}))
		},
		pools() {
			const loggedAddress = this.$store.getters['common/wallet/address']
			if (!loggedAddress) {
				return 0
			}

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
				.filter((pool) => pool.AssetBorrow + pool.BorrowAccrued > 0)
		},
		borrowLimitAsCollateral() {
			return this.assetPoolsWithUser.reduce(
				(acc, userAsset) => (acc += ((((userAsset.AssetDeposit / 1000000) * userAsset.AssetPrice) / 1000000) * userAsset.CollatoralFactor) / 100),
				0
			)
		},
		currentBollowLimit() {
			return (
				this.borrowLimitAsCollateral -
				this.assetPoolsWithUser.reduce((acc, userAsset) => (acc += ((userAsset.AssetBorrow / 1000000) * userAsset.AssetPrice) / 1000000), 0)
			)
		}
	},
	methods: {
		clickAsset(pool) {
			this.$emit('click-asset', pool, 'Borrow')
		},
		getBorrowLimitAssetAmount(pool) {
			const loggedAddress = this.$store.getters['common/wallet/address']
			if (!loggedAddress) {
				return 0
			}
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
			return (
				(assetPools
					.map((pool, index) => ({
						...pool,
						...userAssets[index]
					}))
					.reduce(
						(acc, userAsset) => (acc += ((((userAsset.AssetDeposit / 1000000) * userAsset.AssetPrice) / 1000000) * userAsset.CollatoralFactor) / 100),
						0
					) /
					pool.AssetPrice) *
				1000000
			).toFixed(6)
		}
	}
}
</script>
