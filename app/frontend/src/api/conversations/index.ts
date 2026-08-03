import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { ConversationService } from "#bindings/app/conversation_service";
import type {
	CreateConversationRequest,
	DeleteConversationRequest,
} from "#bindings/app/conversation_service/models";
import type { ValueOf } from "#src/types/helpers";

export const CONVERSATION_QUERY_KEY = {
	LIST_CONVERSATIONS: "ListConversations",
} as const;

export type ConversationQueryKey = ValueOf<typeof CONVERSATION_QUERY_KEY>;

export function useConversations() {
	return useQuery({
		queryKey: [CONVERSATION_QUERY_KEY.LIST_CONVERSATIONS],
		queryFn: () => ConversationService.ListConversations(),
	});
}

export function useCreateConversation() {
	const invalidateConversations = useInvalidateConversations();

	return useMutation({
		mutationFn: (input: CreateConversationRequest) =>
			ConversationService.CreateConversation(input),
		onSuccess: () => {
			invalidateConversations();
		},
	});
}

export function useDeleteConversation() {
	const invalidateConversations = useInvalidateConversations();

	return useMutation({
		mutationFn: (input: DeleteConversationRequest) =>
			ConversationService.DeleteConversation(input),
		onSuccess: () => {
			invalidateConversations();
		},
	});
}

function useInvalidateConversations() {
	const queryClient = useQueryClient();

	return () => {
		queryClient.invalidateQueries({
			queryKey: [CONVERSATION_QUERY_KEY.LIST_CONVERSATIONS],
		});
	};
}
