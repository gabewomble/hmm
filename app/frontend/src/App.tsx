import { MantineProvider } from "@mantine/core";

// core styles are required for all packages
import "@mantine/core/styles.css";
import AppRoutes from "./routes";

// other css files are required only if
// you are using components from the corresponding package
// import '@mantine/dates/styles.css';
// import '@mantine/dropzone/styles.css';
// import '@mantine/code-highlight/styles.css';

function App() {
	return (
		<MantineProvider defaultColorScheme="auto">
			<AppRoutes />
		</MantineProvider>
	);
}

export default App;
