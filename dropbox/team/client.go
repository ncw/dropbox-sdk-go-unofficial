// Copyright (c) Dropbox, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package team : has no documentation (yet)
package team

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/properties"
)

// Client interface describes all routes in this namespace
type Client interface {
	// AlphaGroupsCreate : Creates a new, empty group, with a requested name.
	// Permission : Team member management
	AlphaGroupsCreate(arg *GroupCreateArg) (res *GroupFullInfo, err error)
	// AlphaGroupsGetInfo : Retrieves information about one or more groups.
	// Permission : Team Information
	AlphaGroupsGetInfo(arg *GroupsSelector) (res []*GroupsGetInfoItem, err error)
	// AlphaGroupsList : Lists groups on a team. Permission : Team Information
	AlphaGroupsList(arg *GroupsListArg) (res *GroupsListResult, err error)
	// AlphaGroupsListContinue : Once a cursor has been retrieved from
	// `alphaGroupsList`, use this to paginate through all groups. Permission :
	// Team information
	AlphaGroupsListContinue(arg *GroupsListContinueArg) (res *GroupsListResult, err error)
	// AlphaGroupsUpdate : Updates a group's name, external ID or management
	// type. Permission : Team member management
	AlphaGroupsUpdate(arg *GroupUpdateArgs) (res *GroupFullInfo, err error)
	// DevicesListMemberDevices : List all device sessions of a team's member.
	DevicesListMemberDevices(arg *ListMemberDevicesArg) (res *ListMemberDevicesResult, err error)
	// DevicesListMembersDevices : List all device sessions of a team.
	DevicesListMembersDevices(arg *ListMembersDevicesArg) (res *ListMembersDevicesResult, err error)
	// DevicesListTeamDevices : List all device sessions of a team.
	DevicesListTeamDevices(arg *ListTeamDevicesArg) (res *ListTeamDevicesResult, err error)
	// DevicesRevokeDeviceSession : Revoke a device session of a team's member
	DevicesRevokeDeviceSession(arg *RevokeDeviceSessionArg) (err error)
	// DevicesRevokeDeviceSessionBatch : Revoke a list of device sessions of
	// team members
	DevicesRevokeDeviceSessionBatch(arg *RevokeDeviceSessionBatchArg) (res *RevokeDeviceSessionBatchResult, err error)
	// GetInfo : Retrieves information about a team.
	GetInfo() (res *TeamGetInfoResult, err error)
	// GroupsCreate : Creates a new, empty group, with a requested name.
	// Permission : Team member management
	GroupsCreate(arg *GroupCreateArg) (res *GroupFullInfo, err error)
	// GroupsDelete : Deletes a group. The group is deleted immediately. However
	// the revoking of group-owned resources may take additional time. Use the
	// `groupsJobStatusGet` to determine whether this process has completed.
	// Permission : Team member management
	GroupsDelete(arg *GroupSelector) (res *async.LaunchEmptyResult, err error)
	// GroupsGetInfo : Retrieves information about one or more groups.
	// Permission : Team Information
	GroupsGetInfo(arg *GroupsSelector) (res []*GroupsGetInfoItem, err error)
	// GroupsJobStatusGet : Once an async_job_id is returned from
	// `groupsDelete`, `groupsMembersAdd` , or `groupsMembersRemove` use this
	// method to poll the status of granting/revoking group members' access to
	// group-owned resources. Permission : Team member management
	GroupsJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error)
	// GroupsList : Lists groups on a team. Permission : Team Information
	GroupsList(arg *GroupsListArg) (res *GroupsListResult, err error)
	// GroupsListContinue : Once a cursor has been retrieved from `groupsList`,
	// use this to paginate through all groups. Permission : Team information
	GroupsListContinue(arg *GroupsListContinueArg) (res *GroupsListResult, err error)
	// GroupsMembersAdd : Adds members to a group. The members are added
	// immediately. However the granting of group-owned resources may take
	// additional time. Use the `groupsJobStatusGet` to determine whether this
	// process has completed. Permission : Team member management
	GroupsMembersAdd(arg *GroupMembersAddArg) (res *GroupMembersChangeResult, err error)
	// GroupsMembersList : Lists members of a group. Permission : Team
	// Information
	GroupsMembersList(arg *GroupsMembersListArg) (res *GroupsMembersListResult, err error)
	// GroupsMembersListContinue : Once a cursor has been retrieved from
	// `groupsMembersList`, use this to paginate through all members of the
	// group. Permission : Team information
	GroupsMembersListContinue(arg *GroupsMembersListContinueArg) (res *GroupsMembersListResult, err error)
	// GroupsMembersRemove : Removes members from a group. The members are
	// removed immediately. However the revoking of group-owned resources may
	// take additional time. Use the `groupsJobStatusGet` to determine whether
	// this process has completed. This method permits removing the only owner
	// of a group, even in cases where this is not possible via the web client.
	// Permission : Team member management
	GroupsMembersRemove(arg *GroupMembersRemoveArg) (res *GroupMembersChangeResult, err error)
	// GroupsMembersSetAccessType : Sets a member's access type in a group.
	// Permission : Team member management
	GroupsMembersSetAccessType(arg *GroupMembersSetAccessTypeArg) (res []*GroupsGetInfoItem, err error)
	// GroupsUpdate : Updates a group's name and/or external ID. Permission :
	// Team member management
	GroupsUpdate(arg *GroupUpdateArgs) (res *GroupFullInfo, err error)
	// LinkedAppsListMemberLinkedApps : List all linked applications of the team
	// member. Note, this endpoint does not list any team-linked applications.
	LinkedAppsListMemberLinkedApps(arg *ListMemberAppsArg) (res *ListMemberAppsResult, err error)
	// LinkedAppsListMembersLinkedApps : List all applications linked to the
	// team members' accounts. Note, this endpoint does not list any team-linked
	// applications.
	LinkedAppsListMembersLinkedApps(arg *ListMembersAppsArg) (res *ListMembersAppsResult, err error)
	// LinkedAppsListTeamLinkedApps : List all applications linked to the team
	// members' accounts. Note, this endpoint doesn't list any team-linked
	// applications.
	LinkedAppsListTeamLinkedApps(arg *ListTeamAppsArg) (res *ListTeamAppsResult, err error)
	// LinkedAppsRevokeLinkedApp : Revoke a linked application of the team
	// member
	LinkedAppsRevokeLinkedApp(arg *RevokeLinkedApiAppArg) (err error)
	// LinkedAppsRevokeLinkedAppBatch : Revoke a list of linked applications of
	// the team members
	LinkedAppsRevokeLinkedAppBatch(arg *RevokeLinkedApiAppBatchArg) (res *RevokeLinkedAppBatchResult, err error)
	// MembersAdd : Adds members to a team. Permission : Team member management
	// A maximum of 20 members can be specified in a single call. If no Dropbox
	// account exists with the email address specified, a new Dropbox account
	// will be created with the given email address, and that account will be
	// invited to the team. If a personal Dropbox account exists with the email
	// address specified in the call, this call will create a placeholder
	// Dropbox account for the user on the team and send an email inviting the
	// user to migrate their existing personal account onto the team. Team
	// member management apps are required to set an initial given_name and
	// surname for a user to use in the team invitation and for 'Perform as team
	// member' actions taken on the user before they become 'active'.
	MembersAdd(arg *MembersAddArg) (res *MembersAddLaunch, err error)
	// MembersAddJobStatusGet : Once an async_job_id is returned from
	// `membersAdd` , use this to poll the status of the asynchronous request.
	// Permission : Team member management
	MembersAddJobStatusGet(arg *async.PollArg) (res *MembersAddJobStatus, err error)
	// MembersGetInfo : Returns information about multiple team members.
	// Permission : Team information This endpoint will return
	// `MembersGetInfoItem.id_not_found`, for IDs (or emails) that cannot be
	// matched to a valid team member.
	MembersGetInfo(arg *MembersGetInfoArgs) (res []*MembersGetInfoItem, err error)
	// MembersList : Lists members of a team. Permission : Team information
	MembersList(arg *MembersListArg) (res *MembersListResult, err error)
	// MembersListContinue : Once a cursor has been retrieved from
	// `membersList`, use this to paginate through all team members. Permission
	// : Team information
	MembersListContinue(arg *MembersListContinueArg) (res *MembersListResult, err error)
	// MembersRecover : Recover a deleted member. Permission : Team member
	// management Exactly one of team_member_id, email, or external_id must be
	// provided to identify the user account.
	MembersRecover(arg *MembersRecoverArg) (err error)
	// MembersRemove : Removes a member from a team. Permission : Team member
	// management Exactly one of team_member_id, email, or external_id must be
	// provided to identify the user account. This is not a deactivation where
	// the account can be re-activated again. Calling `membersAdd` with the
	// removed user's email address will create a new account with a new
	// team_member_id that will not have access to any content that was shared
	// with the initial account. This endpoint may initiate an asynchronous job.
	// To obtain the final result of the job, the client should periodically
	// poll `membersRemoveJobStatusGet`.
	MembersRemove(arg *MembersRemoveArg) (res *async.LaunchEmptyResult, err error)
	// MembersRemoveJobStatusGet : Once an async_job_id is returned from
	// `membersRemove` , use this to poll the status of the asynchronous
	// request. Permission : Team member management
	MembersRemoveJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error)
	// MembersSendWelcomeEmail : Sends welcome email to pending team member.
	// Permission : Team member management Exactly one of team_member_id, email,
	// or external_id must be provided to identify the user account. No-op if
	// team member is not pending.
	MembersSendWelcomeEmail(arg *UserSelectorArg) (err error)
	// MembersSetAdminPermissions : Updates a team member's permissions.
	// Permission : Team member management
	MembersSetAdminPermissions(arg *MembersSetPermissionsArg) (res *MembersSetPermissionsResult, err error)
	// MembersSetProfile : Updates a team member's profile. Permission : Team
	// member management
	MembersSetProfile(arg *MembersSetProfileArg) (res *TeamMemberInfo, err error)
	// MembersSuspend : Suspend a member from a team. Permission : Team member
	// management Exactly one of team_member_id, email, or external_id must be
	// provided to identify the user account.
	MembersSuspend(arg *MembersDeactivateArg) (err error)
	// MembersUnsuspend : Unsuspend a member from a team. Permission : Team
	// member management Exactly one of team_member_id, email, or external_id
	// must be provided to identify the user account.
	MembersUnsuspend(arg *MembersUnsuspendArg) (err error)
	// PropertiesTemplateAdd : Add a property template. See route
	// files/properties/add to add properties to a file.
	PropertiesTemplateAdd(arg *AddPropertyTemplateArg) (res *AddPropertyTemplateResult, err error)
	// PropertiesTemplateGet : Get the schema for a specified template.
	PropertiesTemplateGet(arg *properties.GetPropertyTemplateArg) (res *properties.GetPropertyTemplateResult, err error)
	// PropertiesTemplateList : Get the property template identifiers for a
	// team. To get the schema of each template use `propertiesTemplateGet`.
	PropertiesTemplateList() (res *properties.ListPropertyTemplateIds, err error)
	// PropertiesTemplateUpdate : Update a property template. This route can
	// update the template name, the template description and add optional
	// properties to templates.
	PropertiesTemplateUpdate(arg *UpdatePropertyTemplateArg) (res *UpdatePropertyTemplateResult, err error)
	// ReportsGetActivity : Retrieves reporting data about a team's user
	// activity.
	ReportsGetActivity(arg *DateRange) (res *GetActivityReport, err error)
	// ReportsGetDevices : Retrieves reporting data about a team's linked
	// devices.
	ReportsGetDevices(arg *DateRange) (res *GetDevicesReport, err error)
	// ReportsGetMembership : Retrieves reporting data about a team's
	// membership.
	ReportsGetMembership(arg *DateRange) (res *GetMembershipReport, err error)
	// ReportsGetStorage : Retrieves reporting data about a team's storage
	// usage.
	ReportsGetStorage(arg *DateRange) (res *GetStorageReport, err error)
}

type apiImpl dropbox.Context

//AlphaGroupsCreateAPIError is an error-wrapper for the alpha/groups/create route
type AlphaGroupsCreateAPIError struct {
	dropbox.APIError
	EndpointError *GroupCreateError `json:"error"`
}

func (dbx *apiImpl) AlphaGroupsCreate(arg *GroupCreateArg) (res *GroupFullInfo, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "alpha/groups/create"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError AlphaGroupsCreateAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//AlphaGroupsGetInfoAPIError is an error-wrapper for the alpha/groups/get_info route
type AlphaGroupsGetInfoAPIError struct {
	dropbox.APIError
	EndpointError *GroupsGetInfoError `json:"error"`
}

func (dbx *apiImpl) AlphaGroupsGetInfo(arg *GroupsSelector) (res []*GroupsGetInfoItem, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "alpha/groups/get_info"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError AlphaGroupsGetInfoAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//AlphaGroupsListAPIError is an error-wrapper for the alpha/groups/list route
type AlphaGroupsListAPIError struct {
	dropbox.APIError
	EndpointError struct{} `json:"error"`
}

func (dbx *apiImpl) AlphaGroupsList(arg *GroupsListArg) (res *GroupsListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "alpha/groups/list"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError AlphaGroupsListAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//AlphaGroupsListContinueAPIError is an error-wrapper for the alpha/groups/list/continue route
type AlphaGroupsListContinueAPIError struct {
	dropbox.APIError
	EndpointError *GroupsListContinueError `json:"error"`
}

func (dbx *apiImpl) AlphaGroupsListContinue(arg *GroupsListContinueArg) (res *GroupsListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "alpha/groups/list/continue"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError AlphaGroupsListContinueAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//AlphaGroupsUpdateAPIError is an error-wrapper for the alpha/groups/update route
type AlphaGroupsUpdateAPIError struct {
	dropbox.APIError
	EndpointError *GroupUpdateError `json:"error"`
}

func (dbx *apiImpl) AlphaGroupsUpdate(arg *GroupUpdateArgs) (res *GroupFullInfo, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "alpha/groups/update"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError AlphaGroupsUpdateAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//DevicesListMemberDevicesAPIError is an error-wrapper for the devices/list_member_devices route
type DevicesListMemberDevicesAPIError struct {
	dropbox.APIError
	EndpointError *ListMemberDevicesError `json:"error"`
}

func (dbx *apiImpl) DevicesListMemberDevices(arg *ListMemberDevicesArg) (res *ListMemberDevicesResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "devices/list_member_devices"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError DevicesListMemberDevicesAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//DevicesListMembersDevicesAPIError is an error-wrapper for the devices/list_members_devices route
type DevicesListMembersDevicesAPIError struct {
	dropbox.APIError
	EndpointError *ListMembersDevicesError `json:"error"`
}

func (dbx *apiImpl) DevicesListMembersDevices(arg *ListMembersDevicesArg) (res *ListMembersDevicesResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "devices/list_members_devices"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError DevicesListMembersDevicesAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//DevicesListTeamDevicesAPIError is an error-wrapper for the devices/list_team_devices route
type DevicesListTeamDevicesAPIError struct {
	dropbox.APIError
	EndpointError *ListTeamDevicesError `json:"error"`
}

func (dbx *apiImpl) DevicesListTeamDevices(arg *ListTeamDevicesArg) (res *ListTeamDevicesResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "devices/list_team_devices"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError DevicesListTeamDevicesAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//DevicesRevokeDeviceSessionAPIError is an error-wrapper for the devices/revoke_device_session route
type DevicesRevokeDeviceSessionAPIError struct {
	dropbox.APIError
	EndpointError *RevokeDeviceSessionError `json:"error"`
}

func (dbx *apiImpl) DevicesRevokeDeviceSession(arg *RevokeDeviceSessionArg) (err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "devices/revoke_device_session"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError DevicesRevokeDeviceSessionAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	return
}

//DevicesRevokeDeviceSessionBatchAPIError is an error-wrapper for the devices/revoke_device_session_batch route
type DevicesRevokeDeviceSessionBatchAPIError struct {
	dropbox.APIError
	EndpointError *RevokeDeviceSessionBatchError `json:"error"`
}

func (dbx *apiImpl) DevicesRevokeDeviceSessionBatch(arg *RevokeDeviceSessionBatchArg) (res *RevokeDeviceSessionBatchResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "devices/revoke_device_session_batch"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError DevicesRevokeDeviceSessionBatchAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GetInfoAPIError is an error-wrapper for the get_info route
type GetInfoAPIError struct {
	dropbox.APIError
	EndpointError struct{} `json:"error"`
}

func (dbx *apiImpl) GetInfo() (res *TeamGetInfoResult, err error) {
	cli := dbx.Client

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "get_info"), nil)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GetInfoAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsCreateAPIError is an error-wrapper for the groups/create route
type GroupsCreateAPIError struct {
	dropbox.APIError
	EndpointError *GroupCreateError `json:"error"`
}

func (dbx *apiImpl) GroupsCreate(arg *GroupCreateArg) (res *GroupFullInfo, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/create"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsCreateAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsDeleteAPIError is an error-wrapper for the groups/delete route
type GroupsDeleteAPIError struct {
	dropbox.APIError
	EndpointError *GroupDeleteError `json:"error"`
}

func (dbx *apiImpl) GroupsDelete(arg *GroupSelector) (res *async.LaunchEmptyResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/delete"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsDeleteAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsGetInfoAPIError is an error-wrapper for the groups/get_info route
type GroupsGetInfoAPIError struct {
	dropbox.APIError
	EndpointError *GroupsGetInfoError `json:"error"`
}

func (dbx *apiImpl) GroupsGetInfo(arg *GroupsSelector) (res []*GroupsGetInfoItem, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/get_info"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsGetInfoAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsJobStatusGetAPIError is an error-wrapper for the groups/job_status/get route
type GroupsJobStatusGetAPIError struct {
	dropbox.APIError
	EndpointError *GroupsPollError `json:"error"`
}

func (dbx *apiImpl) GroupsJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/job_status/get"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsJobStatusGetAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsListAPIError is an error-wrapper for the groups/list route
type GroupsListAPIError struct {
	dropbox.APIError
	EndpointError struct{} `json:"error"`
}

func (dbx *apiImpl) GroupsList(arg *GroupsListArg) (res *GroupsListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/list"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsListAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsListContinueAPIError is an error-wrapper for the groups/list/continue route
type GroupsListContinueAPIError struct {
	dropbox.APIError
	EndpointError *GroupsListContinueError `json:"error"`
}

func (dbx *apiImpl) GroupsListContinue(arg *GroupsListContinueArg) (res *GroupsListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/list/continue"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsListContinueAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsMembersAddAPIError is an error-wrapper for the groups/members/add route
type GroupsMembersAddAPIError struct {
	dropbox.APIError
	EndpointError *GroupMembersAddError `json:"error"`
}

func (dbx *apiImpl) GroupsMembersAdd(arg *GroupMembersAddArg) (res *GroupMembersChangeResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/members/add"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsMembersAddAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsMembersListAPIError is an error-wrapper for the groups/members/list route
type GroupsMembersListAPIError struct {
	dropbox.APIError
	EndpointError *GroupSelectorError `json:"error"`
}

func (dbx *apiImpl) GroupsMembersList(arg *GroupsMembersListArg) (res *GroupsMembersListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/members/list"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsMembersListAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsMembersListContinueAPIError is an error-wrapper for the groups/members/list/continue route
type GroupsMembersListContinueAPIError struct {
	dropbox.APIError
	EndpointError *GroupsMembersListContinueError `json:"error"`
}

func (dbx *apiImpl) GroupsMembersListContinue(arg *GroupsMembersListContinueArg) (res *GroupsMembersListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/members/list/continue"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsMembersListContinueAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsMembersRemoveAPIError is an error-wrapper for the groups/members/remove route
type GroupsMembersRemoveAPIError struct {
	dropbox.APIError
	EndpointError *GroupMembersRemoveError `json:"error"`
}

func (dbx *apiImpl) GroupsMembersRemove(arg *GroupMembersRemoveArg) (res *GroupMembersChangeResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/members/remove"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsMembersRemoveAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsMembersSetAccessTypeAPIError is an error-wrapper for the groups/members/set_access_type route
type GroupsMembersSetAccessTypeAPIError struct {
	dropbox.APIError
	EndpointError *GroupMemberSetAccessTypeError `json:"error"`
}

func (dbx *apiImpl) GroupsMembersSetAccessType(arg *GroupMembersSetAccessTypeArg) (res []*GroupsGetInfoItem, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/members/set_access_type"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsMembersSetAccessTypeAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//GroupsUpdateAPIError is an error-wrapper for the groups/update route
type GroupsUpdateAPIError struct {
	dropbox.APIError
	EndpointError *GroupUpdateError `json:"error"`
}

func (dbx *apiImpl) GroupsUpdate(arg *GroupUpdateArgs) (res *GroupFullInfo, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "groups/update"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError GroupsUpdateAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//LinkedAppsListMemberLinkedAppsAPIError is an error-wrapper for the linked_apps/list_member_linked_apps route
type LinkedAppsListMemberLinkedAppsAPIError struct {
	dropbox.APIError
	EndpointError *ListMemberAppsError `json:"error"`
}

func (dbx *apiImpl) LinkedAppsListMemberLinkedApps(arg *ListMemberAppsArg) (res *ListMemberAppsResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "linked_apps/list_member_linked_apps"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError LinkedAppsListMemberLinkedAppsAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//LinkedAppsListMembersLinkedAppsAPIError is an error-wrapper for the linked_apps/list_members_linked_apps route
type LinkedAppsListMembersLinkedAppsAPIError struct {
	dropbox.APIError
	EndpointError *ListMembersAppsError `json:"error"`
}

func (dbx *apiImpl) LinkedAppsListMembersLinkedApps(arg *ListMembersAppsArg) (res *ListMembersAppsResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "linked_apps/list_members_linked_apps"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError LinkedAppsListMembersLinkedAppsAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//LinkedAppsListTeamLinkedAppsAPIError is an error-wrapper for the linked_apps/list_team_linked_apps route
type LinkedAppsListTeamLinkedAppsAPIError struct {
	dropbox.APIError
	EndpointError *ListTeamAppsError `json:"error"`
}

func (dbx *apiImpl) LinkedAppsListTeamLinkedApps(arg *ListTeamAppsArg) (res *ListTeamAppsResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "linked_apps/list_team_linked_apps"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError LinkedAppsListTeamLinkedAppsAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//LinkedAppsRevokeLinkedAppAPIError is an error-wrapper for the linked_apps/revoke_linked_app route
type LinkedAppsRevokeLinkedAppAPIError struct {
	dropbox.APIError
	EndpointError *RevokeLinkedAppError `json:"error"`
}

func (dbx *apiImpl) LinkedAppsRevokeLinkedApp(arg *RevokeLinkedApiAppArg) (err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "linked_apps/revoke_linked_app"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError LinkedAppsRevokeLinkedAppAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	return
}

//LinkedAppsRevokeLinkedAppBatchAPIError is an error-wrapper for the linked_apps/revoke_linked_app_batch route
type LinkedAppsRevokeLinkedAppBatchAPIError struct {
	dropbox.APIError
	EndpointError *RevokeLinkedAppBatchError `json:"error"`
}

func (dbx *apiImpl) LinkedAppsRevokeLinkedAppBatch(arg *RevokeLinkedApiAppBatchArg) (res *RevokeLinkedAppBatchResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "linked_apps/revoke_linked_app_batch"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError LinkedAppsRevokeLinkedAppBatchAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersAddAPIError is an error-wrapper for the members/add route
type MembersAddAPIError struct {
	dropbox.APIError
	EndpointError struct{} `json:"error"`
}

func (dbx *apiImpl) MembersAdd(arg *MembersAddArg) (res *MembersAddLaunch, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/add"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersAddAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersAddJobStatusGetAPIError is an error-wrapper for the members/add/job_status/get route
type MembersAddJobStatusGetAPIError struct {
	dropbox.APIError
	EndpointError *async.PollError `json:"error"`
}

func (dbx *apiImpl) MembersAddJobStatusGet(arg *async.PollArg) (res *MembersAddJobStatus, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/add/job_status/get"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersAddJobStatusGetAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersGetInfoAPIError is an error-wrapper for the members/get_info route
type MembersGetInfoAPIError struct {
	dropbox.APIError
	EndpointError *MembersGetInfoError `json:"error"`
}

func (dbx *apiImpl) MembersGetInfo(arg *MembersGetInfoArgs) (res []*MembersGetInfoItem, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/get_info"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersGetInfoAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersListAPIError is an error-wrapper for the members/list route
type MembersListAPIError struct {
	dropbox.APIError
	EndpointError *MembersListError `json:"error"`
}

func (dbx *apiImpl) MembersList(arg *MembersListArg) (res *MembersListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/list"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersListAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersListContinueAPIError is an error-wrapper for the members/list/continue route
type MembersListContinueAPIError struct {
	dropbox.APIError
	EndpointError *MembersListContinueError `json:"error"`
}

func (dbx *apiImpl) MembersListContinue(arg *MembersListContinueArg) (res *MembersListResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/list/continue"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersListContinueAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersRecoverAPIError is an error-wrapper for the members/recover route
type MembersRecoverAPIError struct {
	dropbox.APIError
	EndpointError *MembersRecoverError `json:"error"`
}

func (dbx *apiImpl) MembersRecover(arg *MembersRecoverArg) (err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/recover"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersRecoverAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	return
}

//MembersRemoveAPIError is an error-wrapper for the members/remove route
type MembersRemoveAPIError struct {
	dropbox.APIError
	EndpointError *MembersRemoveError `json:"error"`
}

func (dbx *apiImpl) MembersRemove(arg *MembersRemoveArg) (res *async.LaunchEmptyResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/remove"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersRemoveAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersRemoveJobStatusGetAPIError is an error-wrapper for the members/remove/job_status/get route
type MembersRemoveJobStatusGetAPIError struct {
	dropbox.APIError
	EndpointError *async.PollError `json:"error"`
}

func (dbx *apiImpl) MembersRemoveJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/remove/job_status/get"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersRemoveJobStatusGetAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersSendWelcomeEmailAPIError is an error-wrapper for the members/send_welcome_email route
type MembersSendWelcomeEmailAPIError struct {
	dropbox.APIError
	EndpointError *MembersSendWelcomeError `json:"error"`
}

func (dbx *apiImpl) MembersSendWelcomeEmail(arg *UserSelectorArg) (err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/send_welcome_email"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersSendWelcomeEmailAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	return
}

//MembersSetAdminPermissionsAPIError is an error-wrapper for the members/set_admin_permissions route
type MembersSetAdminPermissionsAPIError struct {
	dropbox.APIError
	EndpointError *MembersSetPermissionsError `json:"error"`
}

func (dbx *apiImpl) MembersSetAdminPermissions(arg *MembersSetPermissionsArg) (res *MembersSetPermissionsResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/set_admin_permissions"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersSetAdminPermissionsAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersSetProfileAPIError is an error-wrapper for the members/set_profile route
type MembersSetProfileAPIError struct {
	dropbox.APIError
	EndpointError *MembersSetProfileError `json:"error"`
}

func (dbx *apiImpl) MembersSetProfile(arg *MembersSetProfileArg) (res *TeamMemberInfo, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/set_profile"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersSetProfileAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//MembersSuspendAPIError is an error-wrapper for the members/suspend route
type MembersSuspendAPIError struct {
	dropbox.APIError
	EndpointError *MembersSuspendError `json:"error"`
}

func (dbx *apiImpl) MembersSuspend(arg *MembersDeactivateArg) (err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/suspend"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersSuspendAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	return
}

//MembersUnsuspendAPIError is an error-wrapper for the members/unsuspend route
type MembersUnsuspendAPIError struct {
	dropbox.APIError
	EndpointError *MembersUnsuspendError `json:"error"`
}

func (dbx *apiImpl) MembersUnsuspend(arg *MembersUnsuspendArg) (err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "members/unsuspend"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError MembersUnsuspendAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	return
}

//PropertiesTemplateAddAPIError is an error-wrapper for the properties/template/add route
type PropertiesTemplateAddAPIError struct {
	dropbox.APIError
	EndpointError *properties.ModifyPropertyTemplateError `json:"error"`
}

func (dbx *apiImpl) PropertiesTemplateAdd(arg *AddPropertyTemplateArg) (res *AddPropertyTemplateResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "properties/template/add"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError PropertiesTemplateAddAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//PropertiesTemplateGetAPIError is an error-wrapper for the properties/template/get route
type PropertiesTemplateGetAPIError struct {
	dropbox.APIError
	EndpointError *properties.PropertyTemplateError `json:"error"`
}

func (dbx *apiImpl) PropertiesTemplateGet(arg *properties.GetPropertyTemplateArg) (res *properties.GetPropertyTemplateResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "properties/template/get"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError PropertiesTemplateGetAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//PropertiesTemplateListAPIError is an error-wrapper for the properties/template/list route
type PropertiesTemplateListAPIError struct {
	dropbox.APIError
	EndpointError *properties.PropertyTemplateError `json:"error"`
}

func (dbx *apiImpl) PropertiesTemplateList() (res *properties.ListPropertyTemplateIds, err error) {
	cli := dbx.Client

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "properties/template/list"), nil)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError PropertiesTemplateListAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//PropertiesTemplateUpdateAPIError is an error-wrapper for the properties/template/update route
type PropertiesTemplateUpdateAPIError struct {
	dropbox.APIError
	EndpointError *properties.ModifyPropertyTemplateError `json:"error"`
}

func (dbx *apiImpl) PropertiesTemplateUpdate(arg *UpdatePropertyTemplateArg) (res *UpdatePropertyTemplateResult, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "properties/template/update"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError PropertiesTemplateUpdateAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//ReportsGetActivityAPIError is an error-wrapper for the reports/get_activity route
type ReportsGetActivityAPIError struct {
	dropbox.APIError
	EndpointError *DateRangeError `json:"error"`
}

func (dbx *apiImpl) ReportsGetActivity(arg *DateRange) (res *GetActivityReport, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "reports/get_activity"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError ReportsGetActivityAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//ReportsGetDevicesAPIError is an error-wrapper for the reports/get_devices route
type ReportsGetDevicesAPIError struct {
	dropbox.APIError
	EndpointError *DateRangeError `json:"error"`
}

func (dbx *apiImpl) ReportsGetDevices(arg *DateRange) (res *GetDevicesReport, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "reports/get_devices"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError ReportsGetDevicesAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//ReportsGetMembershipAPIError is an error-wrapper for the reports/get_membership route
type ReportsGetMembershipAPIError struct {
	dropbox.APIError
	EndpointError *DateRangeError `json:"error"`
}

func (dbx *apiImpl) ReportsGetMembership(arg *DateRange) (res *GetMembershipReport, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "reports/get_membership"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError ReportsGetMembershipAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

//ReportsGetStorageAPIError is an error-wrapper for the reports/get_storage route
type ReportsGetStorageAPIError struct {
	dropbox.APIError
	EndpointError *DateRangeError `json:"error"`
}

func (dbx *apiImpl) ReportsGetStorage(arg *DateRange) (res *GetStorageReport, err error) {
	cli := dbx.Client

	if dbx.Config.Verbose {
		log.Printf("arg: %v", arg)
	}
	b, err := json.Marshal(arg)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", (*dropbox.Context)(dbx).GenerateURL("api", "team", "reports/get_storage"), bytes.NewReader(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if dbx.Config.Verbose {
		log.Printf("req: %v", req)
	}
	resp, err := cli.Do(req)
	if dbx.Config.Verbose {
		log.Printf("resp: %v", resp)
	}
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if dbx.Config.Verbose {
		log.Printf("body: %s", body)
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 409 {
			var apiError ReportsGetStorageAPIError
			err = json.Unmarshal(body, &apiError)
			if err != nil {
				return
			}
			err = apiError
			return
		}
		var apiError dropbox.APIError
		if resp.StatusCode == 400 {
			apiError.ErrorSummary = string(body)
			err = apiError
			return
		}
		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return
		}
		err = apiError
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	return
}

// New returns a Client implementation for this namespace
func New(c dropbox.Config) *apiImpl {
	ctx := apiImpl(dropbox.NewContext(c))
	return &ctx
}
