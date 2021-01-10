//+build wireinject

package autowired

//func AutoWireProcessFileNonRecursive() *service.ProcessFileNonRecursive {
//	wire.Build(service.NewProcessFileNonRecursive)
//	return &service.ProcessFileNonRecursive{}
//}
//
//func AutoWireNewProcessFileUsingRecursiveDepth() *service.ProcessFileUsingRecursiveDepth {
//	wire.Build(service.NewProcessFileUsingRecursiveDepth)
//	return &service.ProcessFileUsingRecursiveDepth{}
//}
//
//func AutoWireNewProcessFileExcludedExtension() *service.ProcessFileExcludedExtension {
//	wire.Build(service.NewProcessFileExcludedExtension)
//	return &service.ProcessFileExcludedExtension{}
//}
//
//func AutoWireNewProcessFileExcludeSpecificFileName() *service.ProcessFileExcludeSpecificFileName {
//	wire.Build(service.NewProcessFileExcludeSpecificFileName)
//	return &service.ProcessFileExcludeSpecificFileName{}
//}
//
//func AutoWireNewProcessFileExcludeFileOfConfigAge() *service.ProcessFileExcludeFileOfConfigAge {
//	wire.Build(service.NewProcessFileExcludeFileOfConfigAge)
//	return &service.ProcessFileExcludeFileOfConfigAge{}
//}
//
//func AutoWireNewProcessFileBackup() *service.ProcessFileBackup {
//	wire.Build(service.NewProcessFileBackup)
//	return &service.ProcessFileBackup{}
//}
//
//func AutoWireNewProcessFileEradicate() *service.ProcessFileEradicate {
//	wire.Build(service.NewProcessFileEradicate)
//	return &service.ProcessFileEradicate{}
//}

////NewFactoryProcessFile
//func AutoWireNewFactoryProcessFile() *service.FactoryProcessFile {
//	wire.Build(service.NewFactoryProcessFile)
//	return &service.FactoryProcessFile{}
//}
//

//func InitializeListener() Listener {
//	wire.Build(
//		NewHelloSpeaker, //*HelloSpeaker, HelloSpeaker doesn't implement Speaker interfaces, conversion is needed in next step
//		wire.Bind(new(Speaker),new(*HelloSpeaker)), //this binding can provide Speaker, not *Speaker
//		NewSimpleListener, //scenario as NewHelloSpeaker
//		//provide Listener, an extra method is needed if *Listener is wanted.
//		wire.Bind(new(Listener), new(*SimpleListener)),
//	)
//	return nil
//}
//
////NewFileJsonParser
//func InitializeListener() service.IFileExpunge {
//	wire.Build(
//		service.NewFileJsonParser, //*HelloSpeaker, HelloSpeaker doesn't implement Speaker interfaces, conversion is needed in next step
//		wire.Bind(new(service.IFileJsonParse),new(*service.FileJsonParser)), //this binding can provide Speaker, not *Speaker
//		service.NewFactoryProcessFile)
//	return nil
//}

