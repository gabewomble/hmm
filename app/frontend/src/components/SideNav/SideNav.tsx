import { Group, NavLink, Stack, Text } from "@mantine/core";
import { Home, MessageSquare, Settings } from "lucide-react";
import { Link, NavLink as RouterNavLink } from "react-router";
import Logo from "../../icons/Logo";
import classes from "./SideNav.module.css";

export default function SideNav() {
	return (
		<Stack className={classes.sideNav}>
			<Link to="/" className={classes.logoLink}>
				<Group className={classes.logo}>
					<Logo size={32} className={classes.logoIcon} />
					<Text fw={700} size="xl">
						hmm
					</Text>
				</Group>
			</Link>

			<NavLink
				label="Home"
				leftSection={<Home size={20} />}
				component={RouterNavLink}
				to="/"
				className={classes.navLink}
			/>

			<NavLink
				label="Messages"
				leftSection={<MessageSquare size={20} />}
				component={RouterNavLink}
				to="/messages"
				className={classes.navLink}
			/>

			<NavLink
				label="Settings"
				leftSection={<Settings size={20} />}
				component={RouterNavLink}
				to="/settings"
				className={classes.navLink}
			/>
		</Stack>
	);
}
