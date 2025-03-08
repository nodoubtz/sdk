/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { accessGroupsListAccessGroupProjects } from "../../funcs/accessGroupsListAccessGroupProjects.js";
import { ListAccessGroupProjectsRequest$inboundSchema } from "../../models/listaccessgroupprojectsop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: ListAccessGroupProjectsRequest$inboundSchema,
};

export const tool$accessGroupsListAccessGroupProjects: ToolDefinition<
  typeof args
> = {
  name: "access-groups_list-access-group-projects",
  description: `List projects of an access group

List projects of an access group`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await accessGroupsListAccessGroupProjects(
      client,
      args.request,
      { fetchOptions: { signal: ctx.signal } },
    ).$inspect();

    if (!result.ok) {
      return {
        content: [{ type: "text", text: result.error.message }],
        isError: true,
      };
    }

    const value = result.value;

    return formatResult(value, apiCall);
  },
};
