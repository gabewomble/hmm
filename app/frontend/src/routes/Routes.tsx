import { Route, Routes } from "react-router";
import Home from "./Home";
import Layout from "./Layout";
import NotFound from "./NotFound";

export default function AppRoutes() {
	return (
		<Routes>
			<Route element={<Layout />}>
				<Route index element={<Home />} />
				<Route path="*" element={<NotFound />} />
			</Route>
		</Routes>
	);
}
