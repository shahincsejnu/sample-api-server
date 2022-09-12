package endpoints

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func GetAllServicesHandler(c *gin.Context) {
	// k8s cluster connection
	k8sClient, err := CreateClient()
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "Could not create the client",
		})
		return
	}

	deployments, err := k8sClient.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "Could not get the deployments!",
		})
		return
	}

	fmt.Println("deployments number : ", len(deployments.Items))

	c.JSON(200, gin.H{
		"Deployments Numbers": len(deployments.Items),
		"Deployment Name":     deployments.Items[0].Name,
		"Replicas":            deployments.Items[0].Spec.Replicas,
	})
}

func CreateClient() (kubernetes.Interface, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kubeConfigPath := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}
