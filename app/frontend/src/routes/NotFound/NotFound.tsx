import { Button, Stack, Text, Title } from "@mantine/core";
import { Link } from "react-router";
import classes from "./NotFound.module.css";

export default function NotFound() {
	return (
		<Stack align="center" justify="center" className={classes.container}>
			<Title order={1}>404</Title>
			<Text>Page not found</Text>
			<Button component={Link} to="/">
				Go home
			</Button>
		</Stack>
	);
}
