import { AppShell } from "@mantine/core";
import { Outlet } from "react-router";
import SideNav from "../components/SideNav";
import classes from "./Layout.module.css";

export default function Layout() {
	return (
		<AppShell navbar={{ width: 250, breakpoint: "sm" }}>
			<AppShell.Navbar>
				<SideNav />
			</AppShell.Navbar>
			<AppShell.Main className={classes.main}>
				<Outlet />
			</AppShell.Main>
		</AppShell>
	);
}
