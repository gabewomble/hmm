import { Group } from "@mantine/core";
import type { ReactNode } from "react";
import classes from "./ContentNavBar.module.css";

interface ContentNavBarProps {
	title: ReactNode;
	actions?: ReactNode;
}

export default function ContentNavBar({ title, actions }: ContentNavBarProps) {
	return (
		<div className={classes.navBar}>
			<Group justify="space-between" w="100%">
				<div className={classes.title}>{title}</div>
				{actions && <div className={classes.actions}>{actions}</div>}
			</Group>
		</div>
	);
}
