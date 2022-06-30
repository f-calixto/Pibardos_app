import { useState } from 'react'
import { Image, Box, AspectRatio, Heading, Center } from 'native-base'
import ViewWithBars from '@Containers/ViewWithBars'
import theme from '@Theme'

const GroupScreen = () => {
  const [isLoading, setIsLoading] = useState(false)

  return (
    <ViewWithBars>
      <Box m={3}>
        <AspectRatio
          ratio={{
            base: 16 / 7
          }}
        >
          <Image
            source={{
              uri: 'https://images.unsplash.com/photo-1543807535-eceef0bc6599?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=387&q=80'
            }}
            alt='group-image'
            borderRadius={10}
          />
        </AspectRatio>
        <Center
          position='absolute'
          bottom={0}
          width='full'
          bgColor='cyan.900'
          opacity='.8'
          height={8}
          borderBottomRadius={10}
        >
          <Heading
            textAlign='center'
            color='white'
          >
            Los Pibardos
          </Heading>
        </Center>
      </Box>
    </ViewWithBars>
  )
}

export default GroupScreen
