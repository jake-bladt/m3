	"github.com/m3db/m3/src/dbnode/encoding/proto"
	"github.com/jhump/protoreflect/desc"
	var schema *desc.MessageDescriptor
	if cfg.Proto != nil {
		logger.Info("Probuf data mode enabled")
		schema, err = proto.ParseProtoSchema(cfg.Proto.SchemaFilePath)
		if err != nil {
			logger.Fatalf("error parsing protobuffer schema: %v", err)
		}
	}

		},
		func(opts client.AdminOptions) client.AdminOptions {
			if cfg.Proto != nil {
				return opts.SetEncodingProto(
					schema,
					encoding.NewOptions(),
				).(client.AdminOptions)
			}
			return opts
		},
	)
	schema *desc.MessageDescriptor,
		if schema != nil {
			enc := proto.NewEncoder(time.Time{}, encodingOpts)
			enc.SetSchema(schema)
			return enc
		}

		return m3tsz.NewEncoder(time.Time{}, nil, nil, m3tsz.DefaultIntOptimizationEnabled, encodingOpts)
		if schema != nil {
			return proto.NewIterator(r, schema, encodingOpts)
		}
		return m3tsz.NewReaderIterator(r, nil, m3tsz.DefaultIntOptimizationEnabled, encodingOpts)
		SetReaderIteratorPool(iteratorPool).