import { Route, Routes } from "react-router";
import ConversationView from "./Conversations/ConversationView";
import NewConversation from "./Conversations/NewConversation";
import Home from "./Home";
import Layout from "./Layout";
import NotFound from "./NotFound";

export default function AppRoutes() {
	return (
		<Routes>
			<Route element={<Layout />}>
				<Route index element={<Home />} />
				<Route path="conversations/new" element={<NewConversation />} />
				<Route path="conversations/:id" element={<ConversationView />} />
				<Route path="*" element={<NotFound />} />
			</Route>
		</Routes>
	);
}
