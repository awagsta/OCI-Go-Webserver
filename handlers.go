package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/awagsta/webserver/helperStructs"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

const compartmentID = "ocid1.compartment.oc1..aaaaaaaa6p62fu3x53o3lwv5p7apjotp7eyfrienrd3xbdy5py7s7pfjgssa"

var config = common.DefaultConfigProvider()

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func listVCN(writer http.ResponseWriter, request *http.Request) {
	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		fmt.Println("Error. Could not authorize user.")
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	req := core.ListVcnsRequest{CompartmentId: common.String(compartmentID)}

	r, err := client.ListVcns(context.Background(), req)

	if err != nil {
		fmt.Println("Error. Could not get VCN List.")
		fmt.Fprintf(writer, "List VCNS failed.")
		return
	}

	var vcnList []helpers.Vcn
	for _, value := range r.Items {
		cidrBlock := value.CidrBlock
		displayName := value.DisplayName
		v := helpers.Vcn{CidrBlock: cidrBlock, DisplayName: displayName}
		vcnList = append(vcnList, v)
	}

	fmt.Println("endpoint hit successfully.")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(vcnList)
}

func listCompute(writer http.ResponseWriter, request *http.Request) {
	client, err := core.NewComputeClientWithConfigurationProvider(config)
	if err != nil {
		fmt.Println("Error. Could not authorize user.")
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	req := core.ListInstancesRequest{CompartmentId: common.String(compartmentID)}

	r, err := client.ListInstances(context.Background(), req)
	if err != nil {
		fmt.Println("Error. Could not get Instance List.")
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	var instanceList []helpers.Instance
	for _, value := range r.Items {
		displayName := value.DisplayName
		shape := value.Shape
		availabilityDomain := value.AvailabilityDomain
		v := helpers.Instance{DisplayName: displayName, Shape: shape, AvailabilityDomain: availabilityDomain}
		instanceList = append(instanceList, v)
	}

	fmt.Println("endpoint hit successfully.")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(instanceList)
}

func createVcn(writer http.ResponseWriter, request *http.Request) {
	var vcnDetails core.CreateVcnDetails

	err := json.NewDecoder(request.Body).Decode(&vcnDetails)
	if err != nil {
		fmt.Println("Error decoding JSON request.")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		fmt.Println("Error. Could not authorize user.")
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	req := core.CreateVcnRequest{}
	req.CidrBlock = vcnDetails.CidrBlock
	req.CompartmentId = vcnDetails.CompartmentId
	req.DisplayName = vcnDetails.DisplayName

	r, err := client.CreateVcn(context.Background(), req)
	if err != nil {
		fmt.Println("Error creating VCN.")
		http.Error(writer, err.Error(), http.StatusForbidden)
		return
	}
	v := helpers.Vcn{CidrBlock: r.CidrBlock, DisplayName: r.DisplayName}
	fmt.Println("VCN Created Successfully")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(v)
}

func deleteVCN(writer http.ResponseWriter, request *http.Request) {
	var vcnID string
	err := json.NewDecoder(request.Body).Decode(&vcnID)
	if err != nil {
		fmt.Println("Error decoding JSON request.")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	c, err := core.NewVirtualNetworkClientWithConfigurationProvider(config)
	if err != nil {
		fmt.Println("Error. Could not authorize user.")
		http.Error(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	req := core.DeleteVcnRequest{VcnId: common.String(vcnID)}

	_, err = c.DeleteVcn(context.Background(), req)
	if err != nil {
		fmt.Println("Delete VCN Failed")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result := "VCN with id " + vcnID + " Deleted Successfully."
	fmt.Println("VCN Deleted Successfully")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)
}
