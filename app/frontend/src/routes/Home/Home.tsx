import { AppShell } from "@mantine/core";
import Header from "../../components/Header";

export default function Home() {
	return (
		<AppShell>
			<AppShell.Header>
				<Header />
			</AppShell.Header>
			<AppShell.Main></AppShell.Main>
			<AppShell.Aside></AppShell.Aside>
		</AppShell>
	);
}
