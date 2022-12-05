package adventofcode;

import java.util.ArrayList;

import org.junit.Test;

public class DayX extends BaseTest {


	public static final int DAY = 0;

	@Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}


	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile(DAY);
		for (final String input : data) {
			System.out.println(input);
		}
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile(DAY);
		for (final String input : data) {
			System.out.println(input);
		}
	}

}
