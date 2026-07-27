import { Divider, NavLink, ScrollArea, Stack, Text } from "@mantine/core";
import { MessageSquarePlus, Settings } from "lucide-react";
import { Link, NavLink as RouterNavLink } from "react-router";
import { useConversations } from "../../api/conversations";
import Logo from "../../icons/Logo";
import classes from "./SideNav.module.css";

export default function SideNav() {
	const { data: conversations, isLoading, isError } = useConversations();

	return (
		<Stack className={classes.sideNav}>
			<Link to="/" className={classes.logoLink}>
				<Stack className={classes.logo} gap={0}>
					<Logo size={32} className={classes.logoIcon} />
					<Text fw={700} size="xl">
						hmm
					</Text>
				</Stack>
			</Link>

			<NavLink
				label="New Conversation"
				leftSection={<MessageSquarePlus size={20} />}
				component={RouterNavLink}
				to="/conversations/new"
				className={classes.navLink}
			/>

			<Divider />

			<ScrollArea className={classes.conversationsScroll}>
				<Stack gap={2}>
					{isLoading && (
						<Text c="dimmed" size="sm">
							Loading...
						</Text>
					)}
					{isError && (
						<Text c="red" size="sm">
							Failed to load conversations
						</Text>
					)}
					{conversations?.map((conv) => (
						<NavLink
							key={conv.id}
							label={conv.name}
							component={RouterNavLink}
							to={`/conversations/${conv.id}`}
							className={classes.navLink}
						/>
					))}
				</Stack>
			</ScrollArea>

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
