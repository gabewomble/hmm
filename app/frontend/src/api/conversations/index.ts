import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { ConversationService } from "#bindings/app/conversation_service";
import type {
	ConversationResponse,
	CreateConversationRequest,
	DeleteConversationRequest,
} from "#bindings/app/conversation_service/models";
import type { ValueOf } from "#src/types/helpers";
import type { IMutationOptions } from "../types";

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

export function useCreateConversation(
	options?: IMutationOptions<ConversationResponse, CreateConversationRequest>,
) {
	const invalidateConversations = useInvalidateConversations();

	return useMutation({
		mutationFn: (input: CreateConversationRequest) =>
			ConversationService.CreateConversation(input),
		onSuccess: (data, variables) => {
			invalidateConversations();
			options?.onSuccess?.(data, variables);
		},
		onError: options?.onError,
	});
}

export function useDeleteConversation(
	options?: IMutationOptions<void, DeleteConversationRequest>,
) {
	const invalidateConversations = useInvalidateConversations();

	return useMutation({
		mutationFn: (input: DeleteConversationRequest) =>
			ConversationService.DeleteConversation(input),
		onSuccess: (data, variables) => {
			invalidateConversations();
			options?.onSuccess?.(data, variables);
		},
		onError: options?.onError,
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
