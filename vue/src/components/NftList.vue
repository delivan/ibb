<template>
	<div class="nft-list">
		<div class="title">NFTs as Collateral</div>
		<div class="all-nfts">
			<div v-if="Array.isArray(userNftList) && userNftList.length > 0" class="sub-title">My NFTs</div>
			<div v-if="Array.isArray(userNftList) && userNftList.length > 0" class="my-nfts">
				<div v-for="nft in userNftList" v-bind:key="nft.id" class="card" @click="clickNftCard(nft)">
					<div class="media">
						<video v-if="nft.ImageUrl.split('.').pop() === 'mp4'" :src="nft.ImageUrl" autoplay="true" loop="true" mute="true" playsinline="true" />
						<img v-else :src="nft.ImageUrl" />
					</div>
					<div class="info">
						<div class="collection">{{ nft.Collection }}</div>
						<div class="name">{{ nft.Name }}</div>
					</div>
					<img v-if="nft.OwnerAddress === 'in escrow'" class="escrow" src="@/assets/images/icons/escrow.png" />
				</div>
			</div>
			<div class="sub-title">NFTs on List</div>
			<div class="nfts-on-list">
				<div v-for="nft in nftList" v-bind:key="nft.id" class="card" @click="clickNftCard(nft)">
					<div class="media">
						<video v-if="nft.ImageUrl.split('.').pop() === 'mp4'" :src="nft.ImageUrl" autoplay="true" loop="true" mute="true" playsinline="true" />
						<img v-else :src="nft.ImageUrl" />
					</div>
					<div class="info">
						<div class="collection">{{ nft.Collection }}</div>
						<div class="name">{{ nft.Name }}</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.nft-list {
	padding: 0 20px;
}

.title {
	font-size: 20px;
}

.sub-title {
	font-size: 20px;
	margin-bottom: 10px;
}

.nfts-in-escrow {
	display: flex;
	flex-wrap: wrap;
	margin-bottom: 8px;
	position: relative;
}

.all-nfts {
	margin-top: 16px;
	padding: 24px 0 0 24px;
	border: 1px solid white;
	border-radius: 10px;
}

.my-nfts {
	margin-top: 16px;
	display: flex;
	flex-wrap: wrap;
	margin-bottom: 8px;
}

.nfts-on-list {
	display: flex;
	flex-wrap: wrap;
}

.escrow {
	position: absolute;
	top: -8px;
	left: -12px;
	width: 224px;
	height: 308px;
}

.card {
	position: relative;
	width: 100%;
	max-width: 200px;
	margin-right: 24px;
	margin-bottom: 24px;
	border: 1px solid rgba(255, 255, 255, 0.6);
	border-radius: 10px;
}

.card:hover {
	cursor: pointer;
}

.media {
	padding: 12px;
	-webkit-box-align: center;
	align-items: center;
	display: flex;
	-webkit-box-pack: center;
	justify-content: center;
	max-height: 100%;
	max-width: 100%;
	height: 200px;
	overflow: hidden;
	position: relative;
	border-radius: 5px;
}

.media > * {
	object-fit: contain;
	width: auto;
	height: auto;
	border-radius: 5px;
	max-width: 100%;
	max-height: 100%;
}

.info {
	padding: 6px 12px 12px;
}

.collection {
	font-size: 13px;
	color: #aaa;
}

.name {
	margin-top: 6px;
	font-size: 16px;
	color: white;
}
</style>

<script>
export default {
	name: 'NftList',
	computed: {
		loggedAddress() {
			return this.$store.getters['common/wallet/address']
		},
		nftListInEscrow() {
			const nftList =
				this.$store.getters['sapienscosmos.ibb.ibb/getNftLoad']({
					params: {
						...(this.loggedAddress && { id: this.loggedAddress })
					}
				})?.DashboardNft ?? []
			return nftList.filter((nft) => nft.OwnerAddress === 'in escrow')
		},
		nftList() {
			const nftList =
				this.$store.getters['sapienscosmos.ibb.ibb/getNftLoad']({
					params: {
						...(this.loggedAddress && { id: this.loggedAddress })
					}
				})?.DashboardNft ?? []
			return nftList.filter((nft) => nft.OwnerAddress !== 'in escrow')
		},
		userNftList() {
			const userNftList =
				this.$store.getters['sapienscosmos.ibb.ibb/getNftLoad']({
					params: {
						id: this.loggedAddress
					}
				})?.UserNft ?? []

			return this.nftListInEscrow.concat(userNftList)
		}
	},
	methods: {
		clickNftCard(nft) {
			console.log(nft)
			this.$emit('click-nft-card', nft)
		}
	}
}
</script>
