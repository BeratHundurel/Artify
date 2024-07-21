<script>
	import '../app.css';
	import { onNavigate } from '$app/navigation';

	onNavigate((navigation) => {
		if (!document.startViewTransition) return;

		return new Promise((resolve) => {
			document.startViewTransition(async () => {
				console.log('startViewTransition');
				resolve();
				await navigation.complete;
			});
		});
	});
</script>

<main>
	<slot></slot>
</main>

<style>
	@keyframes fade-in {
		from {
			opacity: 0;
		}
	}

	@keyframes fade-out {
		to {
			opacity: 0;
		}
	}

	@keyframes slide-from-right {
		from {
			transform: translateX(30px);
		}
	}

	@keyframes slide-to-left {
		to {
			transform: translateX(-30px);
		}
	}

	:root::view-transition-old(root) {
		animation:
			150ms cubic-bezier(0.4, 0, 1, 1) both fade-out,
			450ms cubic-bezier(0.4, 0, 0.2, 1) both slide-to-left;
	}

	:root::view-transition-new(root) {
		animation:
			290ms cubic-bezier(0, 0, 0.2, 1) 120ms both fade-in,
			380ms cubic-bezier(0.4, 0, 0.2, 1) both slide-from-right;
	}
</style>
