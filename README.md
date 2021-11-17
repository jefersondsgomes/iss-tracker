<div align="center">
    <h1>International Space Station Tracker</h1>
    <i>An application that shows the geo localization of ISS in the current time</i>
</div>

</br>

### Technologies:
- [GoLang](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)

</br>

### How to use:
- With GoLang: [clone the project](https://github.com/jefersondsgomes/iss-tracker.git), go to "tracker" folder with any terminal and execute the command "go run main.go";
- With Docker: Pull [iss-tracker](https://hub.docker.com/repository/docker/jefersondsgomes/iss-tracker) image from dockerhub using the command "docker pull jefersondsgomes/iss-tracker", after finish, execute "docker run jefersondsgomes/iss-tracker";
- With Kubernetes (Cluster needed): Navigate to "kubernetes" folder with a terminal and execute "kubectl apply -f .\job.yaml", with this mode, the kubernetes will create a CronJob and execute the application every minute.

</br>

### References:
- [Build Go image with docker](https://docs.docker.com/language/golang/build-images/)
- [CronJobs with kubernetes](https://kubernetes.io/pt-br/docs/concepts/workloads/controllers/cron-jobs/)
- [CronJobs tasks](https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/)