import { AppShell } from "@mantine/core";
import { Outlet } from "react-router";
import SideNav from "../components/SideNav";

export default function Layout() {
	return (
		<AppShell navbar={{ width: 250, breakpoint: "sm" }}>
			<AppShell.Navbar>
				<SideNav />
			</AppShell.Navbar>
			<AppShell.Main>
				<Outlet />
			</AppShell.Main>
		</AppShell>
	);
}
